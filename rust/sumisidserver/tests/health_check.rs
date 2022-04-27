use std::net::TcpListener;
use sqlx::{PgConnection, Connection};

use sumisid::startup::run;
use sumisid::configuration::get_configuration;

pub fn spawn_app() -> String {
    let listener = TcpListener::bind("127.0.0.1:0").expect("Failed to bind random port");
    let port = listener.local_addr().unwrap().port();
    let server = run(listener).expect("Failed to bind address");
    let _ = tokio::spawn(server);

    // return the listening address
    format!("http://127.0.0.1:{}", port)
}

#[tokio::test]
async fn ping_works() {
    // Arrange
    let address = spawn_app();
    let client = reqwest::Client::new();

    // Act
    let response = client
        .get(&format!("{}/ping", &address))
        .send()
        .await
        .expect("Failed to execute http request");

    // Assert
    assert!(response.status().is_success());
    assert_eq!(Some(0), response.content_length());
}


#[tokio::test]
async fn add_club_returns_200_for_valid_form_data() {
    // Arrange
    let address = spawn_app();
    let configuration = get_configuration().expect("Failed to read configuration");
    let connection_string = configuration.database.connection_string();
    // The `Connection` trait MUST be in scope for us to invoke
    // `PgConnection::connect` - it is not an inherent method of the struct!
    let mut connection = PgConnection::connect(&connection_string)
        .await
        .expect("Failed to connect to Postgres.");

    let client = reqwest::Client::new();

    // Act
    let body = "name=Bergen%20Stupeklubb&email=test@test.com";
    let response = client
        .post(&format!("{}/clubs", &address))
        .header("Content-Type", "application/x-www-form-urlencoded")
        .body(body)
        .send()
        .await
        .expect("Failed to execute http request");

    // Assert
    assert_eq!(200, response.status().as_u16());
    let saved = sqlx::query!("SELECT club_name, email FROM clubs",)
        .fetch_one(&mut connection)
        .await
        .expect("Failed to fetch saved subscription.");

    assert_eq!(saved.club_name, "Bergen Stupeklubb");
    assert_eq!(saved.email, Some("test@test.com".to_owned()));
}

#[tokio::test]
async fn add_club_returns_400_for_invalid_form_data() {
    // Arrange
    let address = spawn_app();
    let client = reqwest::Client::new();
    let test_cases = vec![
        ("name=le%20guin", "missing the email"),
        ("email=ursula_le_guin%40gmail.com", "missing the name"),
        ("", "missing both name and email")
    ];

    for(invalid_body, error_message) in test_cases {
        // Act
        let response = client
            .post(&format!("{}/clubs", &address))
            .header("Content-Type", "application/x-www-form-urlencoded")
            .body(invalid_body)
            .send()
            .await
            .expect("Failed to execute http request");
        
        // Assert
        assert_eq!(400, response.status().as_u16(), "Api did not fail (as it should have) with payload {}", error_message);
    }


    // Act
    let body = "name=Bergen%20Stupeklubb&email=1";
    let response = client
        .post(&format!("{}/clubs", &address))
        .header("Content-Type", "application/x-www-form-urlencoded")
        .body(body)
        .send()
        .await
        .expect("Failed to execute http request");

    // Assert
    assert_eq!(200, response.status().as_u16());
}
