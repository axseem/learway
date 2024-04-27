package service_test

import (
	"context"
	"testing"

	"github.com/axseem/learway/model"
	"github.com/axseem/learway/security"
	"github.com/axseem/learway/test"
)

func TestUserUpdatePassword(t *testing.T) {
	testCases := []struct {
		desc  string
		input model.UserUpdatePasswordParams
		err   bool
	}{
		{
			desc: "valid password",
			input: model.UserUpdatePasswordParams{
				CurrentPassword: "0Aaaaaaa",
				NewPassword:     "*UL2}fB!UV&9",
			},
			err: false,
		},
		{
			desc: "invalid password",
			input: model.UserUpdatePasswordParams{
				CurrentPassword: "1Aaaaaaa",
				NewPassword:     "*UL2}fB!UV&9",
			},
			err: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			app, clear := test.App()
			defer clear()

			user, err := app.Service.Session.CreateUser(context.Background(), model.UserCreateParams{
				Username: "user1",
				Email:    "user1@email.com",
				Password: "0Aaaaaaa",
			})
			if err != nil {
				t.Error(err)
			}

			updatedUser, err := app.Service.User.UpdatePassword(context.Background(), user.ID, tC.input)
			if err != nil {
				if tC.err {
					return
				} else {
					t.Error(err)
				}
			}

			err = security.CompareHashAndPassword(updatedUser.Password, tC.input.NewPassword)
			if err != nil {
				t.Error(err)
			}

		})
	}
}

func TestUserUpdateUsername(t *testing.T) {
	testCases := []struct {
		desc  string
		input model.UsernameUpdateParams
		err   bool
	}{
		{
			desc: "valid username",
			input: model.UsernameUpdateParams{
				Username: "newuser1",
			},
			err: false,
		},
		{
			desc: "invalid username",
			input: model.UsernameUpdateParams{
				Username: "newuser 1",
			},
			err: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			app, clear := test.App()
			defer clear()

			user, err := app.Service.Session.CreateUser(context.Background(), model.UserCreateParams{
				Username: "user1",
				Email:    "user1@email.com",
				Password: "0Aaaaaaa",
			})
			if err != nil {
				t.Error(err)
			}

			updatedUser, err := app.Service.User.UpdateUsername(context.Background(), user.ID, tC.input)
			if err != nil {
				if tC.err {
					return
				} else {
					t.Error(err)
				}
			}

			if updatedUser.Username != tC.input.Username {
				t.Error("username has not been updated")
			}
		})
	}
}
