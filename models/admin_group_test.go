package models

import (
	check "gopkg.in/check.v1"
)

func (s *ModelsSuite) TestCreateAdminGroup(c *check.C) {
    admin_group := s.createAdministrationGroupDependencies(c)
    c.Assert(admin_group.Name, check.Equals, "Administrators")
    c.Assert(len(admin_group.Users), check.Equals, 2)

    db_admin_group, err := GetAdminGroup(1)
    c.Assert(err, check.Equals, nil)
    c.Assert(db_admin_group.Name, check.Equals, "Administrators")
    c.Assert(len(db_admin_group.Users), check.Equals, 2)
    c.Assert(db_admin_group.Users[0].Id, check.Equals, int64(1))
    c.Assert(db_admin_group.Users[1].Id, check.Equals, int64(2))

    campaign := s.createCampaignDependencies(c)

    err = PostCampaign(&campaign, campaign.UserId)
    c.Assert(err, check.Equals, nil)

    user_ids, err := GetUsersIDsInUserGroup(db_admin_group.Users[1].Id)
    c.Assert(err, check.Equals, nil)
    c.Assert(len(user_ids), check.Equals, 2)

    campaign, err = GetCampaign(campaign.Id, user_ids)
    campaign_id := campaign.Id

    c.Assert(err, check.Equals, nil)
    c.Assert(campaign.Id, check.Equals, campaign_id)
}

func (s *ModelsSuite) createAdministrationGroupDependencies(ch *check.C) AdminGroup {
    admin_group := AdminGroup{
        Name: "Administrators",
    }

	admin_role, err := GetRoleBySlug(RoleAdmin)
	ch.Assert(err, check.Equals, nil)
	new_admin := User{
		Username: "new-admin",
		Hash:     "123456",
		ApiKey:   "123456",
		Role:     admin_role,
		RoleID:   admin_role.ID,
	}
    user_role, err := GetRoleBySlug(RoleUser)
    ch.Assert(err, check.Equals, nil)
    new_user := User{
        Username: "new-user",
        Hash: "1234567",
        ApiKey: "1234567",
        Role: user_role,
        RoleID: user_role.ID,
    }

    err = PutUser(&new_admin)
    ch.Assert(err, check.Equals, nil)
    err = PutUser(&new_user)
    ch.Assert(err, check.Equals, nil)

    new_admin.Id = 1
    new_user.Id = 2
    users := []User{
        new_admin,
        new_user,
    }
    admin_group.Users = users

    err = PutAdminGroup(&admin_group)
    ch.Assert(err, check.Equals, nil)

    return admin_group
}
