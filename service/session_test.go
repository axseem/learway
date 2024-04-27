package service_test

import (
	"context"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/axseem/learway/model"
	"github.com/axseem/learway/security"
	"github.com/axseem/learway/test"
)

func TestSessionCreateUser(t *testing.T) {
	testCases := []struct {
		desc  string
		input model.UserCreateParams
		err   bool
	}{
		{
			desc: "valid user",
			input: model.UserCreateParams{
				Username: "user1",
				Email:    "user1@email.com",
				Password: "*UL2}fB!UV&9",
			},
			err: false,
		},
		{
			desc: "invalid username",
			input: model.UserCreateParams{
				Username: "user 1",
				Email:    "user1@email.com",
				Password: "*UL2}fB!UV&9",
			},
			err: true,
		},
		{
			desc: "invalid email",
			input: model.UserCreateParams{
				Username: "user1",
				Email:    "user1@email",
				Password: "*UL2}fB!UV&9",
			},
			err: true,
		},
		{
			desc: "invalid password",
			input: model.UserCreateParams{
				Username: "user1",
				Email:    "user1@email.com",
				Password: "00000000",
			},
			err: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			app, clear := test.App()
			defer clear()

			user, err := app.Service.Session.CreateUser(context.Background(), tC.input)
			if err != nil {
				if tC.err {
					return
				} else {
					t.Error(err)
				}
			}

			err = security.CompareHashAndPassword(user.Password, tC.input.Password)
			if err != nil {
				t.Error(err)
			}

			if tC.input.Username != user.Username ||
				tC.input.Email != user.Email {
				t.Errorf("expected: %v, got: %v", tC.input, user)
			}
		})
	}
}

func TestSessionCreate(t *testing.T) {
	testCases := []struct {
		desc  string
		input model.SessionCreateParams
		err   bool
	}{
		{
			desc: "user exist",
			input: model.SessionCreateParams{
				UserID:      "user1",
				Fingerprint: []byte{},
				IP:          net.IP{},
				ExpiresAt:   time.Now().Add(time.Hour),
			},
			err: false,
		},
		{
			desc: "user does not exist",
			input: model.SessionCreateParams{
				UserID:      "user2",
				Fingerprint: []byte{},
				IP:          net.IP{},
				ExpiresAt:   time.Now().Add(time.Hour),
			},
			err: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			app, clear := test.App()
			defer clear()

			_, err := app.Service.Session.CreateUser(context.Background(), model.UserCreateParams{
				Username: "user1",
				Email:    "user1@email.com",
				Password: "*UL2}fB!UV&9",
			})
			if err != nil {
				t.Error(err)
			}

			_, err = app.Service.Session.Create(context.Background(), tC.input)
			if err != nil {
				if tC.err {
					return
				} else {
					t.Error(err)
				}
			}
		})
	}
}

func TestSessionUpdate(t *testing.T) {
	app, clear := test.App()
	defer clear()

	_, err := app.Service.Session.CreateUser(context.Background(), model.UserCreateParams{
		Username: "user1",
		Email:    "user1@email.com",
		Password: "*UL2}fB!UV&9",
	})
	if err != nil {
		t.Error(err)
	}

	session, err := app.Service.Session.Create(context.Background(), model.SessionCreateParams{
		UserID:      "user1",
		Fingerprint: []byte{},
		IP:          net.IP{},
		ExpiresAt:   time.Now().Add(time.Hour),
	})
	if err != nil {
		t.Error(err)
	}

	updatedSession, err := app.Service.Session.Update(context.Background(), session.ID, model.SessionUpdateParams{
		Fingerprint: []byte("Mozilla/5.0"),
		IP:          net.ParseIP("127.0.0.1"),
		ExpiresAt:   time.Now().Add(time.Hour * 2),
	})
	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(session.Fingerprint, updatedSession.Fingerprint) ||
		reflect.DeepEqual(session.IP, updatedSession.IP) ||
		session.ExpiresAt == updatedSession.ExpiresAt {
		t.Error(err)
	}

}
