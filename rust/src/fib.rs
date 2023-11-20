// use std::time::Instant;
// use warp::Filter;
// use reqwest::get;
// use xlsxwriter::*;
// #[tokio::main]
// async fn main() {
//     let route = warp::path!("hello" / String)
//         .map(|name| {
//             format!("Hello, {}!", name)
//         });
 
//     tokio::spawn(async move {
//         warp::serve(route).run(([127, 0, 0, 1], 8080)).await;
//     });
 
//     // Create a new workbook
//     let workbook = Workbook::new("RequestTimes.xlsx");
//     let mut sheet = workbook.add_worksheet(None).unwrap();
 
//     // Send 1000 requests to the server
//     for i in 0..1000 {
//         let start = Instant::now();
//         let resp = get("http://localhost:8080/hello/world").await;
//         match resp {
//             Ok(_) => {
//                 let elapsed = start.elapsed();
//                 // Write the time taken for each request to the Excel sheet
//                 sheet.write_string(i as u32, 0, &format!("{:?}", elapsed), None).unwrap();
//             }
//             Err(err) => println!("Request failed: {:?}", err),
//         }
//     }
 
//     // Save the workbook
//     workbook.close().unwrap();
// }
#[derive(Deserialize)]
struct YourRequestBodyType {
    key: String,
}

use warp::Filter;
use serde::Deserialize;

#[derive(Deserialize)]
struct YourRequestBodyType {
    key: String,
}

#[tokio::main]
async fn main() {
    let get_route = warp::path("your_get_endpoint")
        .and(warp::get())
        .map(|| warp::reply::json(&"Response for GET request"));

    let post_route = warp::path("your_post_endpoint")
        .and(warp::post())
        .and(warp::body::json())
        .map(|body: YourRequestBodyType| {
            warp::reply::json(&format!("Received POST request with body: {:?}", body))
        });

    let routes = get_route.or(post_route);
    warp::serve(routes).run(([127, 0, 0, 1], 3030)).await;
}
