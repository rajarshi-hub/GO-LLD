package main

type userSegment struct {
	userDBConn string
}

func (us *userSegment) createStaticSegment(userIds []string) (string, error) {
	// creating segment and storing in segment Db
	return "1234", nil
}

func (us *userSegment) createDynamicSegment(stringParsedCondition string) (string, error) {
	// creating segment and adding condition
	return "5678", nil
}

func (us *userSegment) getUsersFromSegment(id string) ([]string, error) {
	// Get user segments from id if static else querying the user table if dynamic
	return []string{"12", "34"}, nil
}
