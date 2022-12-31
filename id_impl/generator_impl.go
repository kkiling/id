package generator_id_impl

import (
	"github.com/kkiling/id"
	"github.com/pkg/errors"
	"github.com/sony/sonyflake"
	"time"
)

type service struct {
	flake *sonyflake.Sonyflake
}

func NewService(machineID uint16) (id.IGeneratorId, error) {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{
		MachineID: func() (uint16, error) {
			return machineID, nil
		},
		CheckMachineID: func(_ uint16) bool {
			return true
		},
		StartTime: time.Unix(1703980800, 0),
	})
	return &service{
		flake: flake,
	}, nil
}

func (s *service) GenerateId() id.Uid {
	uid, err := s.flake.NextID()
	if err != nil {
		panic(errors.Wrap(err, "fail generate id"))
	}
	// Сгенерированный id никогда не выйдет за пределы max int64
	return id.Uid(int64(uid))
}
