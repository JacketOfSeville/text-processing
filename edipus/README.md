## Edipus

Edipus is the component responsible for extracting the text content from any file type.
Edipus sends the extracted text to Fausto, via GRPC, so Fausto can store it in the database and call all the plugins.
We have a simple HTML view so you can upload files.
