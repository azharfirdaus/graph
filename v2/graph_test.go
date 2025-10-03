package v2

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

	bandaAcehVertex.AddAdjacency(&medanVertex)
	medanVertex.AddAdjacency(&bandaAcehVertex)
	medanVertex.AddAdjacency(&pekanbaruVertex)
	medanVertex.AddAdjacency(&padangVertex)
	pekanbaruVertex.AddAdjacency(&medanVertex)
	pekanbaruVertex.AddAdjacency(&batamVertex)
	pekanbaruVertex.AddAdjacency(&padangVertex)
	pekanbaruVertex.AddAdjacency(&jambiVertex)
	batamVertex.AddAdjacency(&pekanbaruVertex)
	padangVertex.AddAdjacency(&medanVertex)
	padangVertex.AddAdjacency(&pekanbaruVertex)
	padangVertex.AddAdjacency(&bengkuluVertex)
	jambiVertex.AddAdjacency(&pekanbaruVertex)
	jambiVertex.AddAdjacency(&palembangVertex)
	palembangVertex.AddAdjacency(&jambiVertex)
	palembangVertex.AddAdjacency(&pangkalPinangVertex)
	palembangVertex.AddAdjacency(&bengkuluVertex)
	palembangVertex.AddAdjacency(&bandaLampungVertex)
	pangkalPinangVertex.AddAdjacency(&palembangVertex)
	bengkuluVertex.AddAdjacency(&padangVertex)
	bengkuluVertex.AddAdjacency(&palembangVertex)
	bengkuluVertex.AddAdjacency(&bandaLampungVertex)
	bandaLampungVertex.AddAdjacency(&padangVertex)
	bandaLampungVertex.AddAdjacency(&bengkuluVertex)

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

	bandaAcehVertex.AddAdjacency(&medanVertex)
	medanVertex.AddAdjacency(&bandaAcehVertex)
	medanVertex.AddAdjacency(&pekanbaruVertex)
	medanVertex.AddAdjacency(&padangVertex)
	pekanbaruVertex.AddAdjacency(&medanVertex)
	pekanbaruVertex.AddAdjacency(&batamVertex)
	pekanbaruVertex.AddAdjacency(&padangVertex)
	pekanbaruVertex.AddAdjacency(&jambiVertex)
	batamVertex.AddAdjacency(&pekanbaruVertex)
	padangVertex.AddAdjacency(&medanVertex)
	padangVertex.AddAdjacency(&pekanbaruVertex)
	padangVertex.AddAdjacency(&bengkuluVertex)
	jambiVertex.AddAdjacency(&pekanbaruVertex)
	jambiVertex.AddAdjacency(&palembangVertex)
	palembangVertex.AddAdjacency(&jambiVertex)
	palembangVertex.AddAdjacency(&pangkalPinangVertex)
	palembangVertex.AddAdjacency(&bengkuluVertex)
	palembangVertex.AddAdjacency(&bandaLampungVertex)
	pangkalPinangVertex.AddAdjacency(&palembangVertex)
	bengkuluVertex.AddAdjacency(&padangVertex)
	bengkuluVertex.AddAdjacency(&palembangVertex)
	bengkuluVertex.AddAdjacency(&bandaLampungVertex)
	bandaLampungVertex.AddAdjacency(&padangVertex)
	bandaLampungVertex.AddAdjacency(&bengkuluVertex)

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

	bandaAcehVertex.AddAdjacency(&medanVertex)
	medanVertex.AddAdjacency(&bandaAcehVertex)
	medanVertex.AddAdjacency(&pekanbaruVertex)
	medanVertex.AddAdjacency(&padangVertex)
	pekanbaruVertex.AddAdjacency(&medanVertex)
	pekanbaruVertex.AddAdjacency(&batamVertex)
	pekanbaruVertex.AddAdjacency(&padangVertex)
	pekanbaruVertex.AddAdjacency(&jambiVertex)
	batamVertex.AddAdjacency(&pekanbaruVertex)
	padangVertex.AddAdjacency(&medanVertex)
	padangVertex.AddAdjacency(&pekanbaruVertex)
	padangVertex.AddAdjacency(&bengkuluVertex)
	jambiVertex.AddAdjacency(&pekanbaruVertex)
	jambiVertex.AddAdjacency(&palembangVertex)
	palembangVertex.AddAdjacency(&jambiVertex)
	palembangVertex.AddAdjacency(&pangkalPinangVertex)
	palembangVertex.AddAdjacency(&bengkuluVertex)
	palembangVertex.AddAdjacency(&bandaLampungVertex)
	pangkalPinangVertex.AddAdjacency(&palembangVertex)
	bengkuluVertex.AddAdjacency(&padangVertex)
	bengkuluVertex.AddAdjacency(&palembangVertex)
	bengkuluVertex.AddAdjacency(&bandaLampungVertex)
	bandaLampungVertex.AddAdjacency(&padangVertex)
	bandaLampungVertex.AddAdjacency(&bengkuluVertex)

	vertices := []*Vertex[string]{&bandaAcehVertex, &medanVertex, &pekanbaruVertex,
		&batamVertex, &padangVertex, &jambiVertex, &palembangVertex, &pangkalPinangVertex,
		&bengkuluVertex, &bandaLampungVertex}

	graph, err := Build(len(vertices), vertices)

	assert.Nil(t, err)
	assert.NotNil(t, graph)

	assert.Equal(t, 12, graph.GetEdgeSize())
}
