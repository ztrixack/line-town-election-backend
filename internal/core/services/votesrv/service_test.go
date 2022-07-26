package votesrv_test

import (
	"election-service/internal/core/models"
	"election-service/internal/core/services/votesrv"
	"election-service/internal/repositories/candidatevoterepo"
	"election-service/internal/repositories/electionrepo"
	"election-service/internal/repositories/voterrepo"
	"election-service/internal/sockets/votesock"
	"election-service/internal/utils/resp"
	"election-service/mocks"
	"election-service/pkg"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestService_CreateVoter(t *testing.T) {
	pkg.InitLog()
	assertions := require.New(t)

	type testcase struct {
		name     string
		request  models.CreateVoterData
		expected models.Response
	}

	t.Run("Create Voter", func(t *testing.T) {
		voterRepo := voterrepo.NewMock()
		electionRepo := electionrepo.NewMock()
		candidateVoteRepo := candidatevoterepo.NewMock()
		voteSock := votesock.NewMock()
		voterRepo.On("Create", mocks.NewVoter).Return(&mocks.Voter, nil)
		voterRepo.On("Create", mocks.NilVoter).Return(nil, errors.New(""))
		service := votesrv.New(voterRepo, electionRepo, candidateVoteRepo, voteSock)

		cases := []testcase{
			{name: "Success", request: models.CreateVoterData{CandidateId: 1, NationalId: "1234567890123"}, expected: resp.Created(nil)},
			{name: "Repo error", request: models.CreateVoterData{CandidateId: 0, NationalId: ""}, expected: resp.InternalServerError},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := service.CreateVoter(c.request)
				assertions.Equal(c.expected.Code, actual.Code)
				assertions.Equal(c.expected.Success, actual.Success)
			})
		}
	})
}

func TestService_GetVoter(t *testing.T) {
	pkg.InitLog()
	assertions := require.New(t)

	type testcase struct {
		name     string
		request  string
		expected models.Response
	}

	t.Run("GetVoter", func(t *testing.T) {
		voterRepo := voterrepo.NewMock()
		electionRepo := electionrepo.NewMock()
		candidateVoteRepo := candidatevoterepo.NewMock()
		voteSock := votesock.NewMock()
		voterRepo.On("FindByNationId", "1").Return(&mocks.Voter, nil)
		voterRepo.On("FindByNationId", "2").Return(nil, errors.New(""))
		voterRepo.On("FindByNationId", "3").Return(nil, nil)
		service := votesrv.New(voterRepo, electionRepo, candidateVoteRepo, voteSock)

		cases := []testcase{
			{name: "Success", request: "1", expected: resp.OK(nil)},
			{name: "Repo error", request: "2", expected: resp.InternalServerError},
			{name: "Not Found", request: "3", expected: resp.NotFoundError},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := service.GetVoterByNationId(c.request)
				assertions.Equal(c.expected.Code, actual.Code)
				assertions.Equal(c.expected.Success, actual.Success)
			})
		}
	})
}
