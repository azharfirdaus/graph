package undirected_v2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSize(t *testing.T) {
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

	assert.Equal(t, len(vertices), graph.Size())
}

func TestGetDegreeSize(t *testing.T) {
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

	assert.Equal(t, 24, graph.GetDegreeSize())
}

func TestGetEdgeSize(t *testing.T) {
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

	assert.Equal(t, 12, graph.GetEdgeSize())
}
