package messaging_service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	api "github.com/influenzanet/messaging-service/pkg/api/messaging_service"
	"github.com/influenzanet/messaging-service/pkg/types"
	emailMock "github.com/influenzanet/messaging-service/test/mocks/email-client-service"
	studyMock "github.com/influenzanet/messaging-service/test/mocks/study-service"
	userMock "github.com/influenzanet/messaging-service/test/mocks/user-management-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestSendMessageToAllUsersEndpoint(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockEmailClient := emailMock.NewMockEmailClientServiceApiClient(mockCtrl)
	mockUserClient := userMock.NewMockUserManagementApiClient(mockCtrl)

	s := messagingServer{
		messageDBservice: testMessageDBService,
		clients: &types.APIClients{
			EmailClientService:    mockEmailClient,
			UserManagementService: mockUserClient,
		},
	}

	t.Run("without payload", func(t *testing.T) {
		_, err := s.SendMessageToAllUsers(context.Background(), nil)
		ok, msg := shouldHaveGrpcErrorStatus(err, "missing argument")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with empty payload", func(t *testing.T) {
		_, err := s.SendMessageToAllUsers(context.Background(), &api.SendMessageToAllUsersReq{})
		ok, msg := shouldHaveGrpcErrorStatus(err, "missing argument")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with non admin user", func(t *testing.T) {
		_, err := s.SendMessageToAllUsers(context.Background(), &api.SendMessageToAllUsersReq{
			Token: &api.TokenInfos{
				Id:         "uid",
				InstanceId: testInstanceID,
				Payload: map[string]string{
					"roles":    "PARTICIPANT",
					"username": "testuser",
				},
			},
			Template: &api.EmailTemplate{
				MessageType: "newsletter",
				Translations: []*api.LocalizedTemplate{
					{Lang: "en", Subject: "test", TemplateDef: ""},
				},
			},
		})
		ok, msg := shouldHaveGrpcErrorStatus(err, "not authorized")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with valid user", func(t *testing.T) {
		_, err := s.SendMessageToAllUsers(context.Background(), &api.SendMessageToAllUsersReq{
			Token: &api.TokenInfos{
				Id:         "uid",
				InstanceId: testInstanceID,
				Payload: map[string]string{
					"roles":    "PARTICIPANT,RESEARCHER",
					"username": "testuser",
				},
			},
			Template: &api.EmailTemplate{
				MessageType: "newsletter",
				Translations: []*api.LocalizedTemplate{
					{Lang: "en", Subject: "test", TemplateDef: ""},
				},
			},
		})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

func TestSendMessageToStudyParticipantsEndpoint(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockEmailClient := emailMock.NewMockEmailClientServiceApiClient(mockCtrl)
	mockCtrl2 := gomock.NewController(t)
	defer mockCtrl2.Finish()
	mockUserClient := userMock.NewMockUserManagementApiClient(mockCtrl2)
	mockStudyClient := studyMock.NewMockStudyServiceApiClient(mockCtrl)

	s := messagingServer{
		messageDBservice: testMessageDBService,
		clients: &types.APIClients{
			EmailClientService:    mockEmailClient,
			UserManagementService: mockUserClient,
			StudyService:          mockStudyClient,
		},
	}

	t.Run("without payload", func(t *testing.T) {
		_, err := s.SendMessageToStudyParticipants(context.Background(), nil)
		ok, msg := shouldHaveGrpcErrorStatus(err, "missing argument")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with empty payload", func(t *testing.T) {
		_, err := s.SendMessageToStudyParticipants(context.Background(), &api.SendMessageToStudyParticipantsReq{})
		ok, msg := shouldHaveGrpcErrorStatus(err, "missing argument")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with non admin user", func(t *testing.T) {
		_, err := s.SendMessageToStudyParticipants(context.Background(), &api.SendMessageToStudyParticipantsReq{
			Token: &api.TokenInfos{
				Id:         "uid",
				InstanceId: testInstanceID,
				Payload: map[string]string{
					"roles":    "PARTICIPANT",
					"username": "testuser",
				},
			},
			StudyKey: "testStudy",
			Template: &api.EmailTemplate{
				MessageType: "newsletter",
				Translations: []*api.LocalizedTemplate{
					{Lang: "en", Subject: "test", TemplateDef: ""},
				},
			},
		})
		ok, msg := shouldHaveGrpcErrorStatus(err, "not authorized")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with valid user", func(t *testing.T) {
		_, err := s.SendMessageToStudyParticipants(context.Background(), &api.SendMessageToStudyParticipantsReq{
			Token: &api.TokenInfos{
				Id:         "uid",
				InstanceId: testInstanceID,
				Payload: map[string]string{
					"roles":    "PARTICIPANT,RESEARCHER",
					"username": "testuser",
				},
			},
			Template: &api.EmailTemplate{
				MessageType: "newsletter",
				Translations: []*api.LocalizedTemplate{
					{Lang: "en", Subject: "test", TemplateDef: ""},
				},
			},
		})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

func TestSendInstantEmailEndpoint(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockEmailClient := emailMock.NewMockEmailClientServiceApiClient(mockCtrl)

	s := messagingServer{
		messageDBservice: testMessageDBService,
		clients: &types.APIClients{
			EmailClientService: mockEmailClient,
		},
	}

	_, err := s.messageDBservice.SaveEmailTemplate(testInstanceID, types.EmailTemplate{MessageType: "test-type"})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	t.Run("without payload", func(t *testing.T) {
		_, err := s.SendInstantEmail(context.Background(), nil)
		ok, msg := shouldHaveGrpcErrorStatus(err, "missing argument")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with empty payload", func(t *testing.T) {
		_, err := s.SendInstantEmail(context.Background(), &api.SendEmailReq{})
		ok, msg := shouldHaveGrpcErrorStatus(err, "missing argument")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with not existing template", func(t *testing.T) {
		_, err := s.SendInstantEmail(context.Background(), &api.SendEmailReq{
			InstanceId:  testInstanceID,
			To:          []string{"test@test.test"},
			MessageType: "wrong",
		})
		ok, msg := shouldHaveGrpcErrorStatus(err, "template not found")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with with sending failing", func(t *testing.T) {
		mockEmailClient.EXPECT().SendEmail(
			gomock.Any(),
			gomock.Any(),
		).Return(nil, status.Error(codes.Internal, "failed sending message"))

		_, err := s.SendInstantEmail(context.Background(), &api.SendEmailReq{
			InstanceId:  testInstanceID,
			To:          []string{"test-failed@test.test"},
			MessageType: "test-type",
		})

		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		mails, err := s.messageDBservice.FetchOutgoingEmails(testInstanceID, 1)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
		if len(mails) != 1 || mails[0].To[0] != "test-failed@test.test" {
			t.Errorf("unexpected outgoing mail found: %v", mails)
		}
	})

	t.Run("with with sending succeeded", func(t *testing.T) {
		mockEmailClient.EXPECT().SendEmail(
			gomock.Any(),
			gomock.Any(),
		).Return(nil, nil)

		_, err := s.SendInstantEmail(context.Background(), &api.SendEmailReq{
			InstanceId:  testInstanceID,
			To:          []string{"test-succeeded@test.test"},
			MessageType: "test-type",
		})

		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		mails, err := s.messageDBservice.FetchOutgoingEmails(testInstanceID, 1)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
		if len(mails) != 0 {
			t.Errorf("unexpected outgoing mails found: %v", mails)
		}
	})
}
