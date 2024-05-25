from typing import Dict, List

GRAPH: Dict[str, List[str]]  = {
    "a": ["c"],
    "b": ["c", "d"],
    "c": ["e"],
    "d": ["f"],
    "e": ["f", "h"],
    "f": ["g"],
    "g": [],
    "h": []
}

def topological_sort(entry: str):
    entry_nodes = GRAPH.get(entry, None)

    if entry_nodes is None:
        raise KeyError(f"Node {entry} not in graph")

    seen = set()
    stack = []

    seen.add(entry)
    nodes = [entry, *entry_nodes]

    while len(nodes) > 0:
        node = nodes.pop()

        if node in seen and len(nodes) == 0:
            try:
                not_seen_node = next(filter(lambda key: key not in seen, GRAPH.keys()))
                print("Pick node:", not_seen_node)
                nodes.append(not_seen_node)
            except StopIteration:
                pass

        if node in seen and node not in stack:
            stack.append(node)
            continue

        if node in seen:
            continue

        seen.add(node)
        children = GRAPH.get(node)
        nodes.append(node)
        nodes.extend(children)

    return stack


if __name__ == "__main__":
    print(topological_sort("a"))