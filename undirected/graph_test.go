package undirected

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSize(t *testing.T) {
	bandaAcehVertex := Vertex[string]{"Banda Aceh"}
	medanVertex := Vertex[string]{"Medan"}
	pekanBaruVertex := Vertex[string]{"Pekan Baru"}
	padangVertex := Vertex[string]{"Padang"}
	jambiVertex := Vertex[string]{"Jambi"}
	palembangVertex := Vertex[string]{"Palembang"}
	bengkuluVertex := Vertex[string]{"Bengkulu"}
	pangkalPinangVertex := Vertex[string]{"Pangkal Pinang"}
	bandarLampungVertex := Vertex[string]{"Bandar Lampung"}

	vertices := []*Vertex[string]{&bandaAcehVertex, &medanVertex, &pekanBaruVertex,
		&padangVertex, &jambiVertex, &palembangVertex, &bengkuluVertex,
		&pangkalPinangVertex, &bandarLampungVertex}

	g, err := Build(9, vertices,
		&Edge[string]{Begin: &bandaAcehVertex, End: &medanVertex},
		&Edge[string]{Begin: &medanVertex, End: &bandaAcehVertex},

		&Edge[string]{Begin: &medanVertex, End: &pekanBaruVertex},
		&Edge[string]{Begin: &pekanBaruVertex, End: &medanVertex},

		&Edge[string]{Begin: &medanVertex, End: &padangVertex},
		&Edge[string]{Begin: &padangVertex, End: &medanVertex},

		&Edge[string]{Begin: &pekanBaruVertex, End: &padangVertex},
		&Edge[string]{Begin: &padangVertex, End: &pekanBaruVertex},

		&Edge[string]{Begin: &pekanBaruVertex, End: &jambiVertex},
		&Edge[string]{Begin: &jambiVertex, End: &pekanBaruVertex},

		&Edge[string]{Begin: &padangVertex, End: &bengkuluVertex},
		&Edge[string]{Begin: &bengkuluVertex, End: &padangVertex},

		&Edge[string]{Begin: &jambiVertex, End: &palembangVertex},
		&Edge[string]{Begin: &palembangVertex, End: &jambiVertex},

		&Edge[string]{Begin: &palembangVertex, End: &bengkuluVertex},
		&Edge[string]{Begin: &bengkuluVertex, End: &palembangVertex},

		&Edge[string]{Begin: &palembangVertex, End: &pangkalPinangVertex},
		&Edge[string]{Begin: &pangkalPinangVertex, End: &palembangVertex},

		&Edge[string]{Begin: &palembangVertex, End: &bandarLampungVertex},
		&Edge[string]{Begin: &bandarLampungVertex, End: &palembangVertex},

		&Edge[string]{Begin: &bengkuluVertex, End: &bandarLampungVertex},
		&Edge[string]{Begin: &bandarLampungVertex, End: &bengkuluVertex},
	)

	assert.Nil(t, err)
	assert.NotNil(t, g)

	assert.Equal(t, 9, g.Size())
}

func TestStringInAdjancencyMatrix(t *testing.T) {
	bandaAcehVertex := Vertex[string]{"Banda Aceh"}
	medanVertex := Vertex[string]{"Medan"}
	pekanBaruVertex := Vertex[string]{"Pekan Baru"}
	padangVertex := Vertex[string]{"Padang"}
	jambiVertex := Vertex[string]{"Jambi"}
	palembangVertex := Vertex[string]{"Palembang"}
	bengkuluVertex := Vertex[string]{"Bengkulu"}
	pangkalPinangVertex := Vertex[string]{"Pangkal Pinang"}
	bandarLampungVertex := Vertex[string]{"Bandar Lampung"}

	vertices := []*Vertex[string]{&bandaAcehVertex, &medanVertex, &pekanBaruVertex,
		&padangVertex, &jambiVertex, &palembangVertex, &bengkuluVertex,
		&pangkalPinangVertex, &bandarLampungVertex}

	g, err := Build(9, vertices,
		&Edge[string]{Begin: &bandaAcehVertex, End: &medanVertex},
		&Edge[string]{Begin: &medanVertex, End: &bandaAcehVertex},

		&Edge[string]{Begin: &medanVertex, End: &pekanBaruVertex},
		&Edge[string]{Begin: &pekanBaruVertex, End: &medanVertex},

		&Edge[string]{Begin: &medanVertex, End: &padangVertex},
		&Edge[string]{Begin: &padangVertex, End: &medanVertex},

		&Edge[string]{Begin: &pekanBaruVertex, End: &padangVertex},
		&Edge[string]{Begin: &padangVertex, End: &pekanBaruVertex},

		&Edge[string]{Begin: &pekanBaruVertex, End: &jambiVertex},
		&Edge[string]{Begin: &jambiVertex, End: &pekanBaruVertex},

		&Edge[string]{Begin: &padangVertex, End: &bengkuluVertex},
		&Edge[string]{Begin: &bengkuluVertex, End: &padangVertex},

		&Edge[string]{Begin: &jambiVertex, End: &palembangVertex},
		&Edge[string]{Begin: &palembangVertex, End: &jambiVertex},

		&Edge[string]{Begin: &palembangVertex, End: &bengkuluVertex},
		&Edge[string]{Begin: &bengkuluVertex, End: &palembangVertex},

		&Edge[string]{Begin: &palembangVertex, End: &pangkalPinangVertex},
		&Edge[string]{Begin: &pangkalPinangVertex, End: &palembangVertex},

		&Edge[string]{Begin: &palembangVertex, End: &bandarLampungVertex},
		&Edge[string]{Begin: &bandarLampungVertex, End: &palembangVertex},

		&Edge[string]{Begin: &bengkuluVertex, End: &bandarLampungVertex},
		&Edge[string]{Begin: &bandarLampungVertex, End: &bengkuluVertex},
	)

	assert.Nil(t, err)
	assert.NotNil(t, g)

	adjacencyMatrix := g.GetAdjacencyMatrix()
	assert.True(t, adjacencyMatrix[0][1])
	assert.True(t, adjacencyMatrix[1][0])
	assert.True(t, adjacencyMatrix[1][2])
}
