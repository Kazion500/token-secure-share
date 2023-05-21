<img width="1511" alt="image" src="https://github.com/Kazion500/token-secure-share/assets/64147010/81fe2e58-e0f6-4079-8129-b3d69b73da29">

# Token Secure Share

Token Secure Share is a monorepo project that enables users to securely share token secrets by generating one-time links. The project consists of both frontend and backend components.

## Features
- Generate one-time links: Users can generate unique links that can be used to securely share token secrets. Once the link is accessed, it becomes invalid.
- Secure token sharing: The project ensures that token secrets are shared securely, minimizing the risk of unauthorized access.
- Frontend and backend integration: The frontend and backend components are tightly integrated to provide a seamless user experience.

## Project Structure
The project follows a monorepo structure with the following main components:

`frontend/`: Contains the frontend code for the Token Secure Share application.

`backend/`: Contains the backend code for the Token Secure Share application.

## Installation
To install and run the Token Secure Share project locally, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/your-username/token-secure-share.git
```

2. Install dependencies for both frontend and backend components. In separate terminal windows, 
navigate to the frontend/ and backend/ directories respectively and run the following commands:

```javascript
cd frontend/
npm install

cd ../backend/
npm install
```

3. Configure the project:

   - Frontend: Open the frontend/.env file and provide any necessary configuration variables.
   - Backend: Open the backend/.env file and configure the necessary environment variables, such as database connection details, secret keys, etc.

4. Start the development servers:

   - Frontend: In the frontend/ directory, run:

        ```javascript
        npm start
        ```

    - Backend: In the backend/ directory, run:

        ```javascript
        npm start
        ```

This will start the frontend and backend servers respectively.

5. Access the application:

Open your web browser and navigate to http://localhost:5000 to access the Token Secure Share application.

## Contributing
Contributions to Token Secure Share are welcome! If you want to contribute to this project, please feel free to submit a pull request.

## License
This project is licensed under the MIT License.




