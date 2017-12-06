ALTER TABLE nhn_08_edges
 ADD source BIGINT,
 ADD target BIGINT,
 ADD edge_id BIGINT;

UPDATE nhn_08_edges
SET source = subquery.source,
    target=subquery.target,
    edge_id = subquery.edge_id
FROM (SELECT
       gid,
       ('x'||lpad(from_junct,16,'0'))::BIT(64)::BIGINT AS source,
       ('x'||lpad(to_junct,16,'0'))::BIT(64)::BIGINT AS target,
       ('x'||lpad(nid,16,'0'))::BIT(64)::BIGINT AS edge_id
      FROM nhn_08_edges) AS subquery
WHERE nhn_08_edges.gid=subquery.gid;

-- Lets do the same to the nodes table
ALTER TABLE nhn_08_nodes
 ADD node_id BIGINT;

UPDATE nhn_08_nodes
SET node_id = subquery.node_id
FROM (SELECT
       gid,
       ('x'||lpad(nid,16,'0'))::BIT(64)::BIGINT AS node_id
      FROM nhn_08_nodes) AS subquery
WHERE nhn_08_nodes.gid=subquery.gid;