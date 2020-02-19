package main

import ("fmt"
)

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934


type car struct {
	gas_pedal uint16
	brake_pedal uint16
	steering_wheel int16
	top_speed_kmh float64
}

func (c car) km_h() float64 {
	// c.top_speed_kmh = 500
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}

func (c car) m_h() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax / kmh_multiple)
}

func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed 
}

func newer_top_speed(c car, speed float64) car {
	c.top_speed_kmh = speed
	return c
}



func main() {
	a_car := car{
		gas_pedal: 22341, 
		brake_pedal: 0, 
		steering_wheel: 12651, 
		top_speed_kmh: 225.5}
	// a_car := car{22341,0,12651,225.5}

	fmt.Println("a_car.top_speed_kmh: ", a_car.top_speed_kmh)
	fmt.Println("km_h: ", a_car.km_h())
	fmt.Println("m_h: ", a_car.m_h())

	a_car.new_top_speed(500) 
	fmt.Println("a_car.top_speed_kmh: ", a_car.top_speed_kmh)
	fmt.Println("km_h: ", a_car.km_h())
	fmt.Println("m_h: ", a_car.m_h())

	a_car = newer_top_speed(a_car, 700)
	fmt.Println("a_car.top_speed_kmh: ", a_car.top_speed_kmh)
	fmt.Println("km_h: ", a_car.km_h())
	fmt.Println("m_h: ", a_car.m_h())

	b_car := newer_top_speed(a_car, 700)
	fmt.Println("b_car.top_speed_kmh: ", b_car.top_speed_kmh)
	fmt.Println("km_h: ", b_car.km_h(), "m_h: ", b_car.m_h())

	fmt.Println(a_car)
}