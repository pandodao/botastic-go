package botastic

import "context"

func (s *Suite) TestConversastion() {
	ctx := context.Background()
	conv, err := s.client.CreateConversation(ctx, CreateConversationRequest{
		BotID:        2,
		UserIdentity: "test",
		Lang:         "en",
	})
	s.NoError(err)

	conv, err = s.client.GetConversation(ctx, conv.ID)
	s.NoError(err)

	turn, err := s.client.PostToConversation(ctx, PostToConversationPayloadRequest{
		ConversationID: conv.ID,
		Content:        "test",
		Category:       "plain-text",
	})
	s.NoError(err)

	turn, err = s.client.GetConvTurn(ctx, conv.ID, turn.ID, true)
	s.NoError(err)

	err = s.client.DeleteConversation(ctx, conv.ID)
	s.NoError(err)

	s.T().Log(turn.Response)
}
