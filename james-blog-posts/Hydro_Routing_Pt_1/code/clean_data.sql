-- We need to join the edges to the nodes and delete the nodes that have the same NID and where edges are isolated.
DELETE
FROM nhn_08_nodes N
USING nhn_08_edges E
WHERE (N.nid=E.from_junct OR N.nid=E.to_junct) AND E.isolated = 1;

-- delete edges that are isolated
DELETE
FROM nhn_08_edges
WHERE isolated = 1;