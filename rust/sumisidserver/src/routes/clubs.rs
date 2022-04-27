use actix_web::{web, HttpResponse};

#[derive(serde::Deserialize)]
pub struct FormData {
    name: String,
    email: String,
}

pub async fn create_club(_form: web::Form<FormData>) -> HttpResponse {
    if _form.name.trim().is_empty() || _form.email.is_empty() {
        return HttpResponse::BadRequest().finish()
    }
    HttpResponse::Ok().finish()
}

