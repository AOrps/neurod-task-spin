use std::{
    io::{prelude::*,BufReader},
    net::{TcpListener,TcpStream},
};

fn main() {

    let listener = TcpListener::bind("127.0.0.1:7100").unwrap();    

    for _stream in listener.incoming() {

        let _stream = _stream.unwrap();

        handle_conn(_stream);
        println!("Connection established!");
    }

}

fn handle_conn(mut stream: TcpStream) {
    let buf_reader = BufReader::new(&mut stream);

    let http_request: Vec<_> = buf_reader
        .lines()
        .map(|result| result.unwrap())
        .take_while(|line| !line.is_empty())
        .collect();

    println!("Request: {:#?}", http_request);
}
