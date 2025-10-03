package v2

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

	bandaAcehVertex.AddAdjacency(&medanVertex, 1)
	medanVertex.AddAdjacency(&bandaAcehVertex, 1)
	medanVertex.AddAdjacency(&pekanbaruVertex, 1)
	medanVertex.AddAdjacency(&padangVertex, 1)
	pekanbaruVertex.AddAdjacency(&medanVertex, 1)
	pekanbaruVertex.AddAdjacency(&batamVertex, 1)
	pekanbaruVertex.AddAdjacency(&padangVertex, 1)
	pekanbaruVertex.AddAdjacency(&jambiVertex, 1)
	batamVertex.AddAdjacency(&pekanbaruVertex, 1)
	padangVertex.AddAdjacency(&medanVertex, 1)
	padangVertex.AddAdjacency(&pekanbaruVertex, 1)
	padangVertex.AddAdjacency(&bengkuluVertex, 1)
	jambiVertex.AddAdjacency(&pekanbaruVertex, 1)
	jambiVertex.AddAdjacency(&palembangVertex, 1)
	palembangVertex.AddAdjacency(&jambiVertex, 1)
	palembangVertex.AddAdjacency(&pangkalPinangVertex, 1)
	palembangVertex.AddAdjacency(&bengkuluVertex, 1)
	palembangVertex.AddAdjacency(&bandaLampungVertex, 1)
	pangkalPinangVertex.AddAdjacency(&palembangVertex, 1)
	bengkuluVertex.AddAdjacency(&padangVertex, 1)
	bengkuluVertex.AddAdjacency(&palembangVertex, 1)
	bengkuluVertex.AddAdjacency(&bandaLampungVertex, 1)
	bandaLampungVertex.AddAdjacency(&padangVertex, 1)
	bandaLampungVertex.AddAdjacency(&bengkuluVertex, 1)

	vertices := []*Vertex[string]{&bandaAcehVertex, &medanVertex, &pekanbaruVertex,
		&batamVertex, &padangVertex, &jambiVertex, &palembangVertex, &pangkalPinangVertex,
		&bengkuluVertex, &bandaLampungVertex}

	graph, err := Build(len(vertices), vertices)

	assert.Nil(t, err)
	assert.NotNil(t, graph)

	traversal := StackDFSTraversal[string]{}
	path := traversal.Travel(graph.Size(), &bandaAcehVertex)

	assert.Equal(t, len(vertices), len(path))
	assert.Equal(t, "Banda Aceh", path[0].GetValue())
	assert.Equal(t, "Medan", path[1].GetValue())
	assert.Equal(t, "Padang", path[2].GetValue())
	assert.Equal(t, "Bengkulu", path[3].GetValue())
	assert.Equal(t, "Banda Lampung", path[4].GetValue())
	assert.Equal(t, "Palembang", path[5].GetValue())
	assert.Equal(t, "Pangkal Pinang", path[6].GetValue())
	assert.Equal(t, "Jambi", path[7].GetValue())
	assert.Equal(t, "Pekanbaru", path[8].GetValue())
	assert.Equal(t, "Batam", path[9].GetValue())
}
