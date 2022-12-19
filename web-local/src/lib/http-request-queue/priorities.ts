export const ActivePriority = 10;
export const InactivePriority = 5;

// default is lower. override with higher
// NOTE: should always be less than 100 (reconcile priority)
export const DefaultQueryPriority = 10;
export const QueryPriorities = {
  rows: 50,
  "columns-profile": 45,
  "null-count": 40,
  cardinality: 35,
  "column-cardinality": 35,
  "numeric-histogram": 30,
  topk: 10,
  "rug-histogram": 10,
  "descriptive-statistics": 10,
};