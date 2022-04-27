//! main.rs

use std::net::TcpListener;
use sumisid::startup::run;
use sumisid::configuration::get_configuration;


#[tokio::main]
async fn main() -> std::io::Result<()> {
    // Bubble up the io::Error if we failed to bind the address
    // Otherwise call .await on our Server
    let configuration = get_configuration().expect("Failed to read configuration.");
    let address = format!("127.0.0.1:{}", configuration.application_port);
    println!("Starting server, listening to {}", address);
    let listener = TcpListener::bind(address).expect("Failed to bind to defined address");
    run(listener)?.await
}
