package botastic

import (
	"context"
)

func (s *Suite) TestCreateIndexes() {
	err := s.client.CreateIndexes(context.Background(), CreateIndexesRequest{
		Items: []*CreateIndexesItem{
			{
				ObjectID:   "test-object-id",
				Category:   "plain-text",
				Data:       "test-data",
				Properties: "test-properties",
			},
		},
	})
	s.NoError(err)
}

func (s *Suite) TestDeleteIndex() {
	err := s.client.DeleteIndex(context.Background(), "test-object-id")
	s.NoError(err)
}

func (s *Suite) TestSearchIndexes() {
	_, err := s.client.SearchIndexes(context.Background(), SearchIndexesRequest{
		Keywords: "test",
	})
	s.NoError(err)
}
