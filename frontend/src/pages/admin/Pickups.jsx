import { useEffect, useState } from "react";
import Navbar from "../../components/Navbar";
import BottomNav from "../../components/admin/BottomNav";
import pickupApi from "../../services/pickupApi";

function Pickups() {
  const [pickups, setPickups] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadPickups();
  }, []);

  const loadPickups = async () => {
    try {
      const res = await pickupApi.get("/pickups");
      setPickups(res.data || []);
    } catch (err) {
      console.log(err);
    } finally {
      setLoading(false);
    }
  };

  const handleMarkPickedUp = async (pickupID) => {
    try {
      await pickupApi.patch(`/pickups/${pickupID}/status`, {
        status: "PICKED_UP",
      });
      alert("Pickup ditandai sudah diambil!");
      loadPickups();
    } catch (err) {
      alert("Gagal update status pickup");
    }
  };

  return (
    <>
      <Navbar />
      <div style={{ padding: "20px" }}>
        <h1>Pickup</h1>

        {loading ? (
          <p>Memuat data...</p>
        ) : pickups.length === 0 ? (
          <p>Belum ada pickup.</p>
        ) : (
          <table
            border="1"
            cellPadding="10"
            style={{ width: "100%", borderCollapse: "collapse" }}
          >
            <thead style={{ backgroundColor: "#c0392b", color: "white" }}>
              <tr>
                <th>ID</th>
                <th>No Resi</th>
                <th>User ID</th>
                <th>Berat</th>
                <th>Status Bayar</th>
                <th>Status Pickup</th>
                <th>Dibuat</th>
                <th>Aksi</th>
              </tr>
            </thead>
            <tbody>
              {pickups.map((p) => (
                <tr key={p.pickup_id}>
                  <td>{p.pickup_id}</td>
                  <td>{p.tracking_number}</td>
                  <td>{p.user_id}</td>
                  <td>{p.weight_kg} kg</td>
                  <td>{p.payment_status}</td>
                  <td>
                    <span
                      style={{
                        padding: "4px 8px",
                        borderRadius: "4px",
                        backgroundColor:
                          p.status === "PICKED_UP" ? "#27ae60" : "#e67e22",
                        color: "white",
                        fontSize: "12px",
                      }}
                    >
                      {p.status}
                    </span>
                  </td>
                  <td>
                    {p.created_at
                      ? new Date(p.created_at).toLocaleString("id-ID")
                      : "-"}
                  </td>
                  <td>
                    {p.status === "WAITING_PICKUP" ? (
                      <button
                        onClick={() => handleMarkPickedUp(p.pickup_id)}
                        style={{
                          padding: "4px 8px",
                          backgroundColor: "#c0392b",
                          color: "white",
                          border: "none",
                          borderRadius: "4px",
                          cursor: "pointer",
                        }}
                      >
                        Tandai Diambil
                      </button>
                    ) : (
                      <button
                        disabled
                        style={{
                          padding: "4px 8px",
                          backgroundColor: "#27ae60",
                          color: "white",
                          border: "none",
                          borderRadius: "4px",
                        }}
                      >
                        ✔ Sudah Diambil
                      </button>
                    )}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>
      <BottomNav />
    </>
  );
}

export default Pickups;
