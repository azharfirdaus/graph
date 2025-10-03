package undirected_v2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTravelBfs(t *testing.T) {
	bandaAcehVertex := New("Banda Aceh")
	medanVertex := New("Medan")
	pekanbaruVertex := New("Pekanbaru")
	batamVertex := New("Batam")
	padangVertex := New("Padang")
	jambiVertex := New("Jambi")
	palembangVertex := New("Palembang")
	pangkalPinangVertex := New("Pangkal Pinang")
	bengkuluVertex := New("Bengkulu")
	bandaLampungVertex := New("Banda Lampung")

	bandaAcehVertex.AddNeighbor(&medanVertex)
	medanVertex.AddNeighbor(&bandaAcehVertex)
	medanVertex.AddNeighbor(&pekanbaruVertex)
	medanVertex.AddNeighbor(&padangVertex)
	pekanbaruVertex.AddNeighbor(&medanVertex)
	pekanbaruVertex.AddNeighbor(&batamVertex)
	pekanbaruVertex.AddNeighbor(&padangVertex)
	pekanbaruVertex.AddNeighbor(&jambiVertex)
	batamVertex.AddNeighbor(&pekanbaruVertex)
	padangVertex.AddNeighbor(&medanVertex)
	padangVertex.AddNeighbor(&pekanbaruVertex)
	padangVertex.AddNeighbor(&bengkuluVertex)
	jambiVertex.AddNeighbor(&pekanbaruVertex)
	jambiVertex.AddNeighbor(&palembangVertex)
	palembangVertex.AddNeighbor(&jambiVertex)
	palembangVertex.AddNeighbor(&pangkalPinangVertex)
	palembangVertex.AddNeighbor(&bengkuluVertex)
	palembangVertex.AddNeighbor(&bandaLampungVertex)
	pangkalPinangVertex.AddNeighbor(&palembangVertex)
	bengkuluVertex.AddNeighbor(&padangVertex)
	bengkuluVertex.AddNeighbor(&palembangVertex)
	bengkuluVertex.AddNeighbor(&bandaLampungVertex)
	bandaLampungVertex.AddNeighbor(&padangVertex)
	bandaLampungVertex.AddNeighbor(&bengkuluVertex)

	vertices := []*Vertex[string]{&bandaAcehVertex, &medanVertex, &pekanbaruVertex,
		&batamVertex, &padangVertex, &jambiVertex, &palembangVertex, &pangkalPinangVertex,
		&bengkuluVertex, &bandaLampungVertex}

	graph, err := Build(len(vertices), vertices)

	assert.Nil(t, err)
	assert.NotNil(t, graph)

	result := TravelBfs(graph.Size(), &bandaAcehVertex)

	assert.Equal(t, len(vertices), len(result))
	assert.Equal(t, "Banda Aceh", result[0].GetValue())
	assert.Equal(t, "Medan", result[1].GetValue())
	assert.Equal(t, "Pekanbaru", result[2].GetValue())
	assert.Equal(t, "Padang", result[3].GetValue())
	assert.Equal(t, "Batam", result[4].GetValue())
	assert.Equal(t, "Jambi", result[5].GetValue())
	assert.Equal(t, "Bengkulu", result[6].GetValue())
	assert.Equal(t, "Palembang", result[7].GetValue())
	assert.Equal(t, "Banda Lampung", result[8].GetValue())
	assert.Equal(t, "Pangkal Pinang", result[9].GetValue())

}
