import axios from "axios";

const SudokuAPI = axios.create({
    baseURL: "http://localhost:3001/api",
});

class SudokuService {
    static async findUserByEmail(email: string) {
        try {
            const response = await SudokuAPI.get(`/user/${email}`);

            console.log(response);

            return response.data;
        } catch (error) {
            console.log(error);
        }
    }

    static async registerUser({
        nickname,
        email,
        googleIdToken,
    }: {
        nickname: string;
        email: string;
        googleIdToken: string;
    }) {
        try {
            const response = await SudokuAPI.post("/register", { nickname, email, googleIdToken });

            return response.data;
        } catch (error) {
            console.log(error);
        }
    }
}

export default SudokuService;
