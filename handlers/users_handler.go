package handlers

import (
	"net/http"

	"github.com/showntop/circle-core/logger"
)

type User struct {
}

func (h User) Index(w http.ResponseWriter, r *http.Request) {
	logger.Info("this is the users index page")
	w.Write([]byte("this is the users index page"))
}

func (h User) Show(w http.ResponseWriter, r *http.Request) {
	logger.Info("this is the users show page")
	w.Write([]byte("this is the users show page"))
}

func (h User) Create(w http.ResponseWriter, r *http.Request) {
	logger.Info("this is the users create page")
	w.Write([]byte("this is the users create page"))
}

func (h User) Signup(w http.ResponseWriter, r *http.Request) {
	logger.Info("this is the users create page")
	w.Write([]byte("this is the users create page"))
}

func createUser(c *Context, w http.ResponseWriter, r *http.Request) {
	if !utils.Cfg.EmailSettings.EnableSignUpWithEmail || !utils.Cfg.TeamSettings.EnableUserCreation {
		c.Err = model.NewLocAppError("signupTeam", "api.user.create_user.signup_email_disabled.app_error", nil, "")
		c.Err.StatusCode = http.StatusNotImplemented
		return
	}

	user := model.UserFromJson(r.Body)

	if user == nil {
		c.SetInvalidParam("createUser", "user")
		return
	}

	// the user's username is checked to be valid when they are saved to the database

	user.EmailVerified = false

	var team *model.Team

	if result := <-Srv.Store.Team().Get(user.TeamId); result.Err != nil {
		c.Err = result.Err
		return
	} else {
		team = result.Data.(*model.Team)
	}

	hash := r.URL.Query().Get("h")

	sendWelcomeEmail := true

	if IsVerifyHashRequired(user, team, hash) {
		data := r.URL.Query().Get("d")
		props := model.MapFromJson(strings.NewReader(data))

		if !model.ComparePassword(hash, fmt.Sprintf("%v:%v", data, utils.Cfg.EmailSettings.InviteSalt)) {
			c.Err = model.NewLocAppError("createUser", "api.user.create_user.signup_link_invalid.app_error", nil, "")
			return
		}

		t, err := strconv.ParseInt(props["time"], 10, 64)
		if err != nil || model.GetMillis()-t > 1000*60*60*48 { // 48 hours
			c.Err = model.NewLocAppError("createUser", "api.user.create_user.signup_link_expired.app_error", nil, "")
			return
		}

		if user.TeamId != props["id"] {
			c.Err = model.NewLocAppError("createUser", "api.user.create_user.team_name.app_error", nil, data)
			return
		}

		user.Email = props["email"]
		user.EmailVerified = true
		sendWelcomeEmail = false
	}

	if user.IsSSOUser() {
		user.EmailVerified = true
	}

	if !CheckUserDomain(user, utils.Cfg.TeamSettings.RestrictCreationToDomains) {
		c.Err = model.NewLocAppError("createUser", "api.user.create_user.accepted_domain.app_error", nil, "")
		return
	}

	ruser, err := CreateUser(team, user)
	if err != nil {
		c.Err = err
		return
	}

	if sendWelcomeEmail {
		sendWelcomeEmailAndForget(c, ruser.Id, ruser.Email, team.Name, team.DisplayName, c.GetSiteURL(), c.GetTeamURLFromTeam(team), ruser.EmailVerified)
	}

	w.Write([]byte(ruser.ToJson()))

}
