package undirected_v2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTravelDfs(t *testing.T) {
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

	result := TravelDfs(graph.Size(), &bandaAcehVertex)

	assert.Equal(t, len(vertices), len(result))
	assert.Equal(t, "Banda Aceh", result[0].GetValue())
	assert.Equal(t, "Medan", result[1].GetValue())
	assert.Equal(t, "Padang", result[2].GetValue())
	assert.Equal(t, "Bengkulu", result[3].GetValue())
	assert.Equal(t, "Banda Lampung", result[4].GetValue())
	assert.Equal(t, "Palembang", result[5].GetValue())
	assert.Equal(t, "Pangkal Pinang", result[6].GetValue())
	assert.Equal(t, "Jambi", result[7].GetValue())
	assert.Equal(t, "Pekanbaru", result[8].GetValue())
	assert.Equal(t, "Batam", result[9].GetValue())
}
