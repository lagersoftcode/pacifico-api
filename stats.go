package main

func UpdateUserStats(userId string) {

	var user User
	record := db.Where(User{ID: userId}).First(&user)
	if record.RecordNotFound() {
		panic("user not found at UpdateUserStats: " + userId)
	}
	var totalMedals uint
	db.Model(&ScoreTransaction{}).Where(&ScoreTransaction{UserID: userId, TransactionType: MedalTransaction}).Count(&totalMedals)
	var totalTrophies uint
	db.Model(&ScoreTransaction{}).Where(&ScoreTransaction{UserID: userId, TransactionType: TrophyTransaction}).Count(&totalTrophies)
	var totalKudos uint
	db.Model(&ScoreTransaction{}).Where(&ScoreTransaction{UserID: userId, TransactionType: KudoTransaction}).Count(&totalKudos)

	db.Exec("UPDATE users SET "+
		"stats_total_medals=?, stats_total_trophies=?, stats_total_kudos=?, stats_total_score = "+
		"(SELECT SUM(points) from score_transactions "+
		"WHERE user_id=?)"+
		" WHERE id = ?",
		totalMedals,
		totalTrophies,
		totalKudos,
		userId,
		userId)
}
