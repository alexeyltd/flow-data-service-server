package graph

func (g *DBGraph) FindNode(localId uint) DBNode {
	for _, n := range g.Nodes {
		if n.LocalId == localId {
			return n
		}
	}

	return DBNode{}
}

func (g *DBGraph) FindConnectedNodes(localId uint, slidePort string) []DataConnection {
	if localId == 0 || slidePort == "" {
		return nil
	}

	var nodes []DataConnection

	for _, c := range g.Connections {
		if c.SourceID == localId && c.SourcePort == slidePort {
			nodes = append(nodes, DataConnection{
				SourcePort: c.SourcePort,
				SourceID:   c.SourceID,
				TargetPort: c.TargetPort,
				TargetId:   c.TargetId,
			})
		}
	}

	return nodes
}
