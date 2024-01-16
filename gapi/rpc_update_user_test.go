package gapi

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"testing"
// 	"time"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/require"
// 	mockdb "github.com/valrichter/go-basic-bank/db/mock"
// 	db "github.com/valrichter/go-basic-bank/db/sqlc"
// 	"github.com/valrichter/go-basic-bank/pb"
// 	"github.com/valrichter/go-basic-bank/token"
// 	"github.com/valrichter/go-basic-bank/util"
// 	"google.golang.org/grpc/metadata"
// )

// func TestUpadateUserAPI(t *testing.T) {
// 	user, _ := randomUser(t)

// 	newName := util.RandomOwner()
// 	newEmail := util.RandomEmail()

// 	testCases := []struct {
// 		name          string
// 		req           *pb.UpdateUserRequest
// 		buildStubs    func(store *mockdb.MockStore)
// 		buildContext  func(t *testing.T, tokenMaker token.Maker) context.Context
// 		checkResponse func(t *testing.T, res *pb.UpdateUserResponse, err error)
// 	}{
// 		{
// 			name: "OK",
// 			req: &pb.UpdateUserRequest{
// 				Username: user.Username,
// 				FullName: &newName,
// 				Email:    &newEmail,
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				arg := db.UpdateUserParams{
// 					Username: user.Username,
// 					FullName: sql.NullString{
// 						String: newName,
// 						Valid:  true,
// 					},
// 					Email: sql.NullString{
// 						String: newEmail,
// 						Valid:  true,
// 					},
// 				}
// 				updatedUser := db.User{
// 					Username:         user.Username,
// 					HashedPassword:   user.HashedPassword,
// 					FullName:         newName,
// 					Email:            newEmail,
// 					PasswordChagedAt: user.PasswordChagedAt,
// 					CreatedAt:        user.CreatedAt,
// 					IsEmailVerified:  user.IsEmailVerified,
// 				}
// 				store.EXPECT().
// 					UpdateUser(gomock.Any(), gomock.Eq(arg)).
// 					Times(1).
// 					Return(updatedUser, nil)
// 			},
// 			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
// 				accessToken, _, err := tokenMaker.CreateToken(user.Username, time.Minute)
// 				require.NoError(t, err)
// 				bearerToken := fmt.Sprintf("%s %s", authorizationBearer, accessToken)
// 				md := metadata.MD{
// 					authorizationBearer: []string{
// 						bearerToken,
// 					},
// 				}
// 				return metadata.NewIncomingContext(context.Background(), md)
// 			},
// 			checkResponse: func(t *testing.T, res *pb.UpdateUserResponse, err error) {
// 				require.NoError(t, err)
// 				require.NotNil(t, res)
// 				updatedUser := res.GetUser()
// 				require.Equal(t, user.Username, updatedUser.Username)
// 				require.Equal(t, newName, updatedUser.FullName)
// 				require.Equal(t, newEmail, updatedUser.Email)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			storeCtrl := gomock.NewController(t)
// 			defer storeCtrl.Finish()
// 			store := mockdb.NewMockStore(storeCtrl)

// 			tc.buildStubs(store)
// 			server := newTestServer(t, store, nil)

// 			ctx := tc.buildContext(t, server.tokenMaker)
// 			res, err := server.UpdateUser(ctx, tc.req)
// 			tc.checkResponse(t, res, err)
// 		})
// 	}
// }
