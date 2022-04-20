package request

type TaskReq struct {
	Device   string `json:"device" form:"device" validate:"required"`
	App      string `json:"app" form:"app" validate:"required"`
	Duration string `json:"duration" form:"duration" validate:"required"`
	Report   string `json:"report" form:"report" validate:"required"`
	CleanLog *bool  `json:"cleanLog" form:"cleanLog" validate:"required"`
	UserId   string `json:"userId" form:"userId" validate:"required"`
}