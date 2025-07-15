// Builder pattern for creating a graph structure in Rust.

use std::error::Error;

struct Graph {
    nodes: Vec<String>,
    edges: Vec<(String, String)>,
}

struct GraphBuilder {
    nodes: Vec<String>,
    edges: Vec<(String, String)>,
}

impl GraphBuilder {
    fn new() -> Self {
        GraphBuilder {
            nodes: Vec::new(),
            edges: Vec::new(),
        }
    }

    fn with_nodes(mut self, nodes: Vec<String>) -> Self {
        self.nodes = nodes;
        self
    }

    fn with_edges(mut self, edges: Vec<(String, String)>) -> Self {
        self.edges = edges;
        self
    }

    fn add_node(mut self, node: String) -> Self {
        self.nodes.push(node);
        self
    }

    fn add_edge(mut self, from: String, to: String) -> Self {
        self.edges.push((from, to));
        self
    }

    fn build(self) -> Result<Graph, Box<dyn Error>> {
        if self.nodes.is_empty() {
            return Err("Graph must have at least one node".into());
        }
        if self.edges.is_empty() && self.nodes.len() > 1 {
            return Err("Graph must have at least one edge if there are multiple nodes".into());
        }
        Ok(Graph {
            nodes: self.nodes,
            edges: self.edges,
        })
    }
}

impl Graph {
    fn builder() -> GraphBuilder {
        GraphBuilder::new()
    }
}

fn render_dot(graph: &Graph) -> String {
    let mut dot = String::from("digraph G {\n");
    for node in &graph.nodes {
        dot.push_str(&format!("  {};\n", node));
    }
    for (from, to) in &graph.edges {
        dot.push_str(&format!("  {} -> {};\n", from, to));
    }
    dot.push_str("}\n");
    dot
}

fn main() -> Result<(), Box<dyn Error>> {
    let graph = Graph::builder()
        .with_nodes(vec!["A".to_string(), "B".to_string(), "C".to_string()])
        .with_edges(vec![
            ("A".to_string(), "B".to_string()),
            ("B".to_string(), "C".to_string()),
        ])
        .build()?;

    let dot_representation = render_dot(&graph);
    println!("{}", dot_representation);

    let graph2 = Graph::builder()
        .add_node("X".to_string())
        .add_node("Y".to_string())
        .add_node("Z".to_string())
        .add_edge("X".to_string(), "Y".to_string())
        .add_edge("Y".to_string(), "Z".to_string())
        .build()?;

    println!("\nsecond graph:");
    println!("{}", render_dot(&graph2));

    println!("\ntrying to build an invalid graph (no nodes):");
    match Graph::builder().build() {
        Ok(_) => println!("graph built successfully"),
        Err(e) => println!("error: {}", e),
    }

    println!("\ntrying to build an invalid graph (multiple nodes, no edges):");
    match Graph::builder()
        .add_node("A".to_string())
        .add_node("B".to_string())
        .build() {
        Ok(_) => println!("graph built successfully"),
        Err(e) => println!("error: {}", e),
    }

    Ok(())
}