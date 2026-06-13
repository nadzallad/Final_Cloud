import api from "./api";

export const payOrder = (data) => {
    return api.post("/payments", data);
};