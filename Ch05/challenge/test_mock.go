package main

const mock1 = `
[
	{
		"start": "2021-01-01T13:32",
		"end": "2021-01-01T13:40",
		"id": "car1",
		"distance": 1.7
	  },
	  {
	"start": "2021-01-01T13:32",
	"end": "2021-01-01T13:40",
	"id": "car1",
	"distance": 1.71
	  }
]
`

const mock2 = `
[
	{"start": "2021-01-01T13:32", "end": "2021-01-01T13:40", "id": "car1", "distance": 1.7},
	{"start": "2021-01-01T13:37", "end": "2021-01-01T13:41", "id": "car2", "distance": 2.7},
	{"start": "2021-01-01T18:11", "end": "2021-01-01T19:03", "id": "car1", "distance": 2.2},
	{"start": "2021-01-02T03:23", "end": "2021-01-02T03:40", "id": "car1", "distance": 1.1},
	{"start": "2021-01-02T06:07", "end": "2021-01-02T06:14", "id": "car1", "distance": 0.7},
	{"start": "2021-01-02T08:51", "end": "2021-01-02T09:12", "id": "car2", "distance": 12.3},
	{"start": "2021-01-03T12:18", "end": "2021-01-03T12:29", "id": "car2", "distance": 0.3},
	{"start": "2021-01-03T22:17", "end": "2021-01-03T23:02", "id": "car1", "distance": 12.13}
]
`
