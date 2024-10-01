var a InfoMessage
	a.Distance = w.distance()
	a.TrainingType = w.TrainingType
	a.Duration = w.Duration
	a.Speed = w.meanSpeed()
	a.Calories = w.Calories()
	return a