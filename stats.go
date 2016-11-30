package main

func UpdateUserStats(userId string) {

	var userStatus UserStatus
	record := db.Where(&UserStatus{UserId: userId}).First(&userStatus)
	if record.RecordNotFound() {
		userStatus = UserStatus{
			UserId: userId,
		}
		db.Create(&userStatus)
	}
	var totalMedals uint
	db.Model(&ScoreTransaction{}).Where(&ScoreTransaction{UserID: userId, TransactionType: MedalTransaction}).Count(&totalMedals)
	var totalTrophies uint
	db.Model(&ScoreTransaction{}).Where(&ScoreTransaction{UserID: userId, TransactionType: TrophyTransaction}).Count(&totalTrophies)
	var totalKudos uint
	db.Model(&ScoreTransaction{}).Where(&ScoreTransaction{UserID: userId, TransactionType: KudoTransaction}).Count(&totalKudos)

	db.Exec("UPDATE user_statuses SET "+
		"total_medals=?, total_trophies=?, total_kudos=?, total_score = "+
		"(SELECT SUM(points) from score_transactions "+
		"WHERE user_id=?)"+
		" WHERE user_id = ?",
		totalMedals,
		totalTrophies,
		totalKudos,
		userId,
		userId)
}
