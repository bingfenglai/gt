package service

import (
	"github.com/bingfenglai/gt/common/errors"
	"github.com/bingfenglai/gt/convert"
	"github.com/bingfenglai/gt/domain/dto"
	"github.com/bingfenglai/gt/storage"
)

type roleServiceImpl struct {
}

func (svc *roleServiceImpl) GetSessionRolesByUid(uid string) ([]*dto.RoleSessionDTO, error) {

	if uid == "" {
		return nil, errors.ErrUserIDCannotBeEmpty
	}

	roles, err := storage.RoleStorage.GetRolesByUid(uid)

	return convert.Role2RoleSessionDTOList(roles), err
}
