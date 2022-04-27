//! startup.rs
use actix_web::{web, App, HttpServer};
use actix_web::dev::Server;
use std::net::TcpListener;

use crate::routes::{ping, create_club};

pub fn run(listener: TcpListener) -> Result<Server, std::io::Error>  {
    let server = HttpServer::new(|| {
        App::new()
            .route("/ping", web::get().to(ping))
            .route("/clubs", web::post().to(create_club))
    })
    .listen(listener)?
    .run();
    Ok(server)
}
