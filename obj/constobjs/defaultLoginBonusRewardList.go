package constobjs

import "github.com/RunnersRevival/outrun/obj"

var DefaultLoginBonusRewardList = func() []obj.LoginBonusReward {
	return []obj.LoginBonusReward{
		// Day 1
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("910000", 3000), // 3000 Rings
						obj.NewItem("240000", 1),    // 1 Item Roulette Ticket
					},
				),
			},
		),
		// Day 2
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("910000", 3000), // 3000 Rings
						obj.NewItem("240000", 1),    // 1 Item Roulette Ticket
					},
				),
			},
		),
		// Day 3
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("910000", 5000), // 5000 Rings
						obj.NewItem("240000", 1),    // 1 Item Roulette Ticket
					},
				),
			},
		),
		// Day 4
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("910000", 5000), // 5000 Rings
						obj.NewItem("240000", 1),    // 1 Item Roulette Ticket
					},
				),
			},
		),
		// Day 5
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("900000", 10),   // 10 Red Star Rings
						obj.NewItem("910000", 5000), // 5000 Rings
						obj.NewItem("240000", 2),    // 2 Item Roulette Tickets
					},
				),
			},
		),
		// Day 6
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("910000", 10000), // 10000 Rings
						obj.NewItem("240000", 2),     // 2 Item Roulette Tickets
					},
				),
			},
		),
		// Day 7
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("900000", 20),    // 20 Red Star Rings
						obj.NewItem("910000", 15000), // 15000 Rings
						obj.NewItem("240000", 2),     // 2 Item Roulette Tickets
					},
				),
			},
		),
	}
}()

// TODO: Remove this when finishing troubleshooting the login bonus!
var TestLoginBonusRewardList = func() []obj.LoginBonusReward {
	return []obj.LoginBonusReward{
		// Day 1
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("910000", 1000), // 1000 Rings
						obj.NewItem("900000", 1),    // 1 Red Ring
					},
				),
			},
		),
		// Day 2
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("910000", 2000), // 2000 Rings
						obj.NewItem("900000", 2),    // 2 Red Rings
					},
				),
			},
		),
		// Day 3
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("910000", 3000), // 3000 Rings
						obj.NewItem("900000", 3),    // 3 Red Rings
					},
				),
			},
		),
		// Day 4
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("910000", 4000), // 4000 Rings
						obj.NewItem("900000", 4),    // 4 Red Rings
					},
				),
			},
		),
		// Day 5
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("910000", 5000), // 5000 Rings
						obj.NewItem("900000", 5),    // 5 Red Rings
					},
				),
			},
		),
		// Day 6
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("910000", 6000), // 6000 Rings
						obj.NewItem("900000", 6),    // 6 Red Rings
					},
				),
			},
		),
		// Day 7
		obj.NewLoginBonusReward(
			[]obj.SelectReward{
				obj.NewSelectReward(
					[]obj.Item{
						obj.NewItem("910000", 7000), // 7000 Rings
						obj.NewItem("900000", 7),    // 7 Red Rings
					},
				),
			},
		),
	}
}()
