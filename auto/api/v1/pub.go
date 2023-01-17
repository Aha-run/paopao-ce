// Code generated by go-mir. DO NOT EDIT.
// versions:
// - mir v3.0.1

package v1

import (
	"net/http"

	"github.com/alimy/mir/v3"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

type Pub interface {
	TopicList(*web.TopicListReq) (*web.TopicListResp, mir.Error)
	TweetComments(*web.TweetCommentsReq) (*web.TweetCommentsResp, mir.Error)
	TweetDetail(*web.TweetDetailReq) (*web.TweetDetailResp, mir.Error)
	SendCaptcha(*web.SendCaptchaReq) mir.Error
	GetCaptcha() (*web.GetCaptchaResp, mir.Error)
	Register(*web.RegisterReq) (*web.RegisterResp, mir.Error)
	Login(*web.LoginReq) (*web.LoginResp, mir.Error)
	Version() (*web.VersionResp, mir.Error)

	mustEmbedUnimplementedPubServant()
}

type PubBinding interface {
	BindTopicList(*gin.Context) (*web.TopicListReq, mir.Error)
	BindTweetComments(*gin.Context) (*web.TweetCommentsReq, mir.Error)
	BindTweetDetail(*gin.Context) (*web.TweetDetailReq, mir.Error)
	BindSendCaptcha(*gin.Context) (*web.SendCaptchaReq, mir.Error)
	BindRegister(*gin.Context) (*web.RegisterReq, mir.Error)
	BindLogin(*gin.Context) (*web.LoginReq, mir.Error)

	mustEmbedUnimplementedPubBinding()
}

type PubRender interface {
	RenderTopicList(*gin.Context, *web.TopicListResp, mir.Error)
	RenderTweetComments(*gin.Context, *web.TweetCommentsResp, mir.Error)
	RenderTweetDetail(*gin.Context, *web.TweetDetailResp, mir.Error)
	RenderSendCaptcha(*gin.Context, mir.Error)
	RenderGetCaptcha(*gin.Context, *web.GetCaptchaResp, mir.Error)
	RenderRegister(*gin.Context, *web.RegisterResp, mir.Error)
	RenderLogin(*gin.Context, *web.LoginResp, mir.Error)
	RenderVersion(*gin.Context, *web.VersionResp, mir.Error)

	mustEmbedUnimplementedPubRender()
}

// RegisterPubServant register Pub servant to gin
func RegisterPubServant(e *gin.Engine, s Pub, b PubBinding, r PubRender) {
	router := e.Group("v1")

	// register routes info to router
	router.Handle("GET", "/tags", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindTopicList(c)
		if err != nil {
			r.RenderTopicList(c, nil, err)
			return
		}
		resp, err := s.TopicList(req)
		r.RenderTopicList(c, resp, err)
	})

	router.Handle("GET", "/post/comments", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindTweetComments(c)
		if err != nil {
			r.RenderTweetComments(c, nil, err)
			return
		}
		resp, err := s.TweetComments(req)
		r.RenderTweetComments(c, resp, err)
	})

	router.Handle("GET", "/post", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindTweetDetail(c)
		if err != nil {
			r.RenderTweetDetail(c, nil, err)
			return
		}
		resp, err := s.TweetDetail(req)
		r.RenderTweetDetail(c, resp, err)
	})

	router.Handle("POST", "/captcha", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindSendCaptcha(c)
		if err != nil {
			r.RenderSendCaptcha(c, err)
			return
		}
		r.RenderSendCaptcha(c, s.SendCaptcha(req))
	})

	router.Handle("GET", "/captcha", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		resp, err := s.GetCaptcha()
		r.RenderGetCaptcha(c, resp, err)
	})

	router.Handle("POST", "/auth/register", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindRegister(c)
		if err != nil {
			r.RenderRegister(c, nil, err)
			return
		}
		resp, err := s.Register(req)
		r.RenderRegister(c, resp, err)
	})

	router.Handle("POST", "/auth/login", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindLogin(c)
		if err != nil {
			r.RenderLogin(c, nil, err)
			return
		}
		resp, err := s.Login(req)
		r.RenderLogin(c, resp, err)
	})

	router.Handle("GET", "/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		resp, err := s.Version()
		r.RenderVersion(c, resp, err)
	})

}

// UnimplementedPubServant can be embedded to have forward compatible implementations.
type UnimplementedPubServant struct {
}

func (UnimplementedPubServant) TopicList(req *web.TopicListReq) (*web.TopicListResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPubServant) TweetComments(req *web.TweetCommentsReq) (*web.TweetCommentsResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPubServant) TweetDetail(req *web.TweetDetailReq) (*web.TweetDetailResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPubServant) SendCaptcha(req *web.SendCaptchaReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPubServant) GetCaptcha() (*web.GetCaptchaResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPubServant) Register(req *web.RegisterReq) (*web.RegisterResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPubServant) Login(req *web.LoginReq) (*web.LoginResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPubServant) Version() (*web.VersionResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedPubServant) mustEmbedUnimplementedPubServant() {}

// UnimplementedPubRender can be embedded to have forward compatible implementations.
type UnimplementedPubRender struct {
	RenderAny func(*gin.Context, any, mir.Error)
}

func (r *UnimplementedPubRender) RenderTopicList(c *gin.Context, data *web.TopicListResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedPubRender) RenderTweetComments(c *gin.Context, data *web.TweetCommentsResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedPubRender) RenderTweetDetail(c *gin.Context, data *web.TweetDetailResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedPubRender) RenderSendCaptcha(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedPubRender) RenderGetCaptcha(c *gin.Context, data *web.GetCaptchaResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedPubRender) RenderRegister(c *gin.Context, data *web.RegisterResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedPubRender) RenderLogin(c *gin.Context, data *web.LoginResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedPubRender) RenderVersion(c *gin.Context, data *web.VersionResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedPubRender) mustEmbedUnimplementedPubRender() {}

// UnimplementedPubBinding can be embedded to have forward compatible implementations.
type UnimplementedPubBinding struct {
	BindAny func(*gin.Context, any) mir.Error
}

func (b *UnimplementedPubBinding) BindTopicList(c *gin.Context) (*web.TopicListReq, mir.Error) {
	obj := new(web.TopicListReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPubBinding) BindTweetComments(c *gin.Context) (*web.TweetCommentsReq, mir.Error) {
	obj := new(web.TweetCommentsReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPubBinding) BindTweetDetail(c *gin.Context) (*web.TweetDetailReq, mir.Error) {
	obj := new(web.TweetDetailReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPubBinding) BindSendCaptcha(c *gin.Context) (*web.SendCaptchaReq, mir.Error) {
	obj := new(web.SendCaptchaReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPubBinding) BindRegister(c *gin.Context) (*web.RegisterReq, mir.Error) {
	obj := new(web.RegisterReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPubBinding) BindLogin(c *gin.Context) (*web.LoginReq, mir.Error) {
	obj := new(web.LoginReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedPubBinding) mustEmbedUnimplementedPubBinding() {}
