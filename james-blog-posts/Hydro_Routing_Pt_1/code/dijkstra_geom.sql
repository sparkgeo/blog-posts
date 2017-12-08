SELECT node, edge, cost, agg_cost, direction, source, target, geom
FROM (SELECT *
      FROM pgr_dijkstra(
                  edges_sql :='SELECT edge_id AS id,' ||
                                'source, ' ||
                                'target, ' ||
                                'st_length(geom)::DOUBLE PRECISION AS cost' ||
                              'FROM nhn_08_edges' ,
                  start_vid := -1262225606265190345,
                  end_vid   := 5509843120790651269
              )
     ) AS route
  JOIN nhn_08_edges
    ON route.edge = nhn_08_edges.edge_id