import { useEffect, useState } from "react";
import { getAllShipments } from "../../services/shipment";
import Navbar from "../../components/Navbar";

function Shipment() {
  const [shipments, setShipments] = useState([]);

  useEffect(() => {
    loadShipments();
  }, []);

  const loadShipments = async () => {
    try {
      const res = await getAllShipments();
      setShipments(res.data);
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <>
      <Navbar />

      <div style={{ padding: "20px" }}>
        <h1>Shipments</h1>

        {shipments.length === 0 ? (
          <p>Belum ada shipment.</p>
        ) : (
          <table
            border="1"
            cellPadding="10"
            style={{
              width: "100%",
              borderCollapse: "collapse",
            }}
          >
            <thead
              style={{
                backgroundColor: "#c0392b",
                color: "white",
              }}
            >
              <tr>
                <th>ID</th>
                <th>Tracking ID</th>
                <th>No Resi</th>
                <th>Asal</th>
                <th>Tujuan</th>
                <th>Lokasi Saat Ini</th>
                <th>Status</th>
                <th>ETA</th>
              </tr>
            </thead>

            <tbody>
              {shipments.map((s) => (
                <tr key={s.shipment_id}>
                  <td>{s.shipment_id}</td>
                  <td>{s.tracking_id}</td>
                  <td>{s.no_resi}</td>
                  <td>{s.origin_city}</td>
                  <td>{s.destination_city}</td>
                  <td>{s.current_location}</td>

                  <td>
                    <span
                      style={{
                        padding: "4px 8px",
                        borderRadius: "4px",
                        backgroundColor:
                          s.status === "DELIVERED"
                            ? "#27ae60"
                            : "#e67e22",
                        color: "white",
                        fontSize: "12px",
                      }}
                    >
                      {s.status}
                    </span>
                  </td>

                  <td>
                    {s.eta
                      ? new Date(s.eta).toLocaleString("id-ID")
                      : "-"}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>
    </>
  );
}

export default Shipment;
