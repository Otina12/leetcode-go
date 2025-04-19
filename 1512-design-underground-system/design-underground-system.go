type StationTimeTuple struct {
    station string
    time int
}

type StationsTuple struct {
    station1 string
    station2 string
}

type UndergroundSystem struct {
    CustomerLastCheckIn map[int]StationTimeTuple
    StationTimes map[StationsTuple][]int
}


func Constructor() UndergroundSystem {
    return UndergroundSystem { make(map[int]StationTimeTuple), make(map[StationsTuple][]int) }
}


func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
    this.CustomerLastCheckIn[id] = StationTimeTuple { stationName, t }
}


func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
    lastCheckIn := this.CustomerLastCheckIn[id]
    stationsTuple := StationsTuple { lastCheckIn.station, stationName }
    
    this.StationTimes[stationsTuple] = append(this.StationTimes[stationsTuple], t - lastCheckIn.time)
}


func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
    stationsTuple := StationsTuple { startStation, endStation }
    stationTimes := this.StationTimes[stationsTuple]
    total := 0

    for _, time := range stationTimes {
        total += time
    }

    return float64(total) / float64(len(stationTimes))
}


/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */