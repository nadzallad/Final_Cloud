import { useState } from "react";
import trackingService from "../services/trackingService";
import "./TrackingUser.css";

function TrackingUser() {
  const [noResi, setNoResi] = useState("");
  const [trackingData, setTrackingData] = useState([]);
  const [loading, setLoading] = useState(false);

  const handleSearch = async () => {
    try {
      setLoading(true);

      const response = await trackingService.getTrackingByResi(noResi);

      setTrackingData(response.data);
    } catch (error) {
      console.error(error);
      alert("Data tracking tidak ditemukan");
      setTrackingData([]);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="tracking-user-container">
      <h1>📦 Lacak Paket</h1>

      <div className="search-box">
        <input
          type="text"
          placeholder="Masukkan Nomor Resi"
          value={noResi}
          onChange={(e) => setNoResi(e.target.value)}
        />

        <button onClick={handleSearch}>
          Lacak
        </button>
      </div>

      {loading && <p>Loading...</p>}

      {trackingData.length > 0 && (
        <div className="tracking-result">
          <h3>Riwayat Pengiriman</h3>

          {trackingData.map((item) => (
            <div key={item.tracking_id} className="tracking-item">
              <p>
                <strong>Status:</strong> {item.status}
              </p>

              <p>
                <strong>Lokasi:</strong> {item.location}
              </p>

              <p>
                <strong>Catatan:</strong> {item.note}
              </p>

              <p>
                <strong>Waktu:</strong> {item.created_at}
              </p>

              <hr />
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default TrackingUser;