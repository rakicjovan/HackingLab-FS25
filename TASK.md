# Task: Insecure File Upload Vulnerability

## Story

Your colleague built a simple web app to upload and share pictures. As you know, he’s a bit of a "vibe coder" — fast and loose with best practices, you offered to do a code review to improve security, but he didn’t take it well. Instead, he challenged you: he added a secret `/flag` route somewhere in the backend, and it's your job to find it.

He used PHP on the backend, so you suspect you might be able to execute PHP code through the file upload functionality.

## What You Know

- The app only allows uploading **PNG** and **JPEG** image files.
- There is client-side validation that tries to enforce this, but client-side checks can be bypassed.
- The uploaded files are accessible under `/uploads/`.
- There is a hidden `/flag` route on the backend, but visiting it via the browser only returns the default `404 Not Found` message by NGINX.

## Your Task

- Start the service and visit it using the given FQDN.
- Investigate how to bypass the image upload restrictions.
- Try to upload a PHP script disguised as an image or bypass the file type checks.
- Use your uploaded PHP script to access the secret `/flag` route and retrieve the flag.
- Remember: this is a learning exercise to understand why client-side validation is not enough and why backend security is crucial.

## Tips

- Use browser developer tools to modify the HTML or the form data before submitting your upload.
- Try manipulating the file extension or content-type headers.
- Remember that the backend might still execute PHP code if your upload bypasses validation.
- Think about how PHP execution works in common web servers and how you might trigger it.

Good luck — and happy hacking!  
