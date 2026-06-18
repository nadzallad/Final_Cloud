import { useEffect, useState } from "react";
import { getAllShipments, getWarehouseByResi, getShipmentByResi, getDeliveryByResi } from "../../services/shipmentTrackingService";
import Navbar from "../../components/Navbar";

// ── Helper: susun timeline dari data warehouse + shipment + delivery ──────────
function buildTimeline(warehouse, shipment, delivery) {
  const steps = [];

  // Step 1: Masuk gudang asal (dari warehouse-service)
  if (warehouse) {
    steps.push({
      icon: "📦",
      title: "Paket Diterima di Gudang",
      lokasi: "Gudang Asal",
      waktu: warehouse.created_at
        ? new Date(warehouse.created_at).toLocaleString("id-ID") : "-",
      status: "IN_WAREHOUSE",
      done: true,
      current: !shipment && !delivery && warehouse.status === "IN_WAREHOUSE",
    });
  }

  // Step 2: Keluar gudang / siap dikirim
  if (warehouse?.status === "OUT_FOR_SHIPMENT" || shipment || delivery) {
    steps.push({
      icon: "🏭",
      title: "Keluar dari Gudang",
      lokasi: shipment?.origin_city ?? "Gudang Asal",
      waktu: "-",
      status: "OUT_FOR_SHIPMENT",
      done: true,
      current: false,
    });
  }

  // Step 3: Shipment (perjalanan antar kota / antar gudang)
  if (shipment) {
    const isCurrentStep = !delivery;
    steps.push({
      icon: "🚛",
      title: "Dalam Perjalanan Antar Kota",
      lokasi: `${shipment.origin_city ?? "-"} → ${shipment.destination_city ?? "-"}`,
      waktu: shipment.eta
        ? `ETA: ${new Date(shipment.eta).toLocaleString("id-ID")}` : "-",
      status: shipment.status,
      done: true,
      current: isCurrentStep,
      currentLocation: shipment.current_location,
    });
  }

  // Step 4: Delivery (kurir ke penerima)
  if (delivery) {
    const delivered = delivery.status === "DELIVERED";
    steps.push({
      icon: delivered ? "✅" : "🚚",
      title: delivered ? "Paket Diterima Penerima" : "Dikirim ke Penerima",
      lokasi: delivery.delivery_address ?? "-",
      waktu: delivered && delivery.delivered_at
        ? new Date(delivery.delivered_at).toLocaleString("id-ID")
        : "Sedang dalam pengiriman",
      status: delivery.status,
      done: delivered,
      current: !delivered,
      currentLocation: delivery.delivery_address,
    });
  }

  return steps;
}

// ── Status badge ───────────────────────────────────────────────────────────────
const STATUS_MAP = {
  IN_WAREHOUSE:     { bg: "#2563eb", label: "Di Gudang" },
  OUT_FOR_SHIPMENT: { bg: "#7c3aed", label: "Keluar Gudang" },
  IN_TRANSIT:       { bg: "#d97706", label: "Dalam Perjalanan" },
  OUT_FOR_DELIVERY: { bg: "#ea580c", label: "Dikirim Kurir" },
  DELIVERED:        { bg: "#16a34a", label: "Terkirim" },
};

function Badge({ status }) {
  const s = STATUS_MAP[status] ?? { bg: "#6b7280", label: status };
  return (
    <span style={{
      background: s.bg, color: "white",
      padding: "2px 10px", borderRadius: 20,
      fontSize: 11, fontWeight: 700,
    }}>
      {s.label}
    </span>
  );
}

// ── Komponen utama ─────────────────────────────────────────────────────────────
export default function Shipment() {
  const [shipments, setShipments] = useState([]);

  const [noResi,   setNoResi]   = useState("");
  const [result,   setResult]   = useState(null);   // { warehouse, shipment, delivery } | { notFound }
  const [loading,  setLoading]  = useState(false);
  const [searched, setSearched] = useState("");

  useEffect(() => { loadAll(); }, []);

  const loadAll = async () => {
    try {
      const data = await getAllShipments();
      setShipments(Array.isArray(data) ? data : data?.data ?? []);
    } catch (_) {}
  };

  const handleCari = async () => {
    const resi = noResi.trim();
    if (!resi) return;
    setLoading(true);
    setResult(null);
    setSearched(resi);

    let warehouse = null, shipment = null, delivery = null;

    try { const r = await getWarehouseByResi(resi); if (r?.found) warehouse = r.data; } catch (_) {}
    try { shipment = await getShipmentByResi(resi); }  catch (_) {}
    try { delivery = await getDeliveryByResi(resi); }  catch (_) {}

    if (!warehouse && !shipment && !delivery) {
      setResult({ notFound: true });
    } else {
      setResult({ warehouse, shipment, delivery });
    }

    setLoading(false);
  };

  const timeline = result && !result.notFound
    ? buildTimeline(result.warehouse, result.shipment, result.delivery)
    : [];

  // Cari step yang sedang aktif (posisi paket sekarang)
  const currentStep = timeline.findLast?.(s => s.current) ?? timeline[timeline.length - 1];

  return (
    <>
      <Navbar />
      <div style={{ padding: "24px", maxWidth: 920, margin: "0 auto" }}>
        <h1 style={{ marginBottom: 4, color: "#1a1a2e" }}>📦 Shipments</h1>
        <p style={{ color: "#888", fontSize: 13, marginBottom: 24 }}>
          Kelola dan lacak perjalanan paket berdasarkan nomor resi.
        </p>

        {/* ── Panel Lacak ── */}
        <div style={{
          background: "white", borderRadius: 14, padding: "20px 24px",
          marginBottom: 28, boxShadow: "0 2px 10px rgba(0,0,0,0.07)",
          border: "1px solid #e5e7eb",
        }}>
          <h3 style={{ margin: "0 0 4px", fontSize: 15, color: "#1a1a2e" }}>
            🔍 Lacak Paket via No Resi
          </h3>
          <p style={{ margin: "0 0 14px", fontSize: 12, color: "#888" }}>
            Masukkan nomor resi untuk melihat paket ini sudah melewati gudang mana saja dan sekarang ada di mana.
          </p>

          {/* Search box */}
          <div style={{ display: "flex", gap: 10 }}>
            <input
              type="text"
              placeholder="Contoh: TRK-20240001"
              value={noResi}
              onChange={e => { setNoResi(e.target.value); setResult(null); }}
              onKeyDown={e => e.key === "Enter" && handleCari()}
              style={{
                flex: 1, padding: "11px 14px",
                border: "1px solid #d1d5db", borderRadius: 8,
                fontSize: 14, outline: "none",
              }}
            />
            <button
              onClick={handleCari}
              disabled={loading}
              style={{
                padding: "11px 24px",
                background: loading ? "#93c5fd" : "#2563eb",
                color: "white", border: "none", borderRadius: 8,
                fontWeight: 700, cursor: loading ? "not-allowed" : "pointer",
                fontSize: 14,
              }}
            >
              {loading ? "Mencari…" : "Cari"}
            </button>
          </div>

          {/* Not found */}
          {result?.notFound && (
            <div style={{
              marginTop: 14, padding: "12px 16px",
              background: "#fef2f2", border: "1px solid #fecaca", borderRadius: 10,
            }}>
              <p style={{ margin: 0, color: "#dc2626", fontWeight: 700 }}>
                ❌ Paket tidak ditemukan
              </p>
              <p style={{ margin: "4px 0 0", fontSize: 12, color: "#6b7280" }}>
                No resi <strong>{searched}</strong> belum terdaftar di sistem.
              </p>
            </div>
          )}

          {/* Hasil */}
          {timeline.length > 0 && (
            <div style={{ marginTop: 20 }}>

              {/* Posisi sekarang — highlight */}
              {currentStep && (
                <div style={{
                  background: "linear-gradient(135deg, #1d4ed8, #2563eb)",
                  borderRadius: 12, padding: "14px 18px",
                  marginBottom: 20, color: "white",
                }}>
                  <p style={{ margin: "0 0 4px", fontSize: 12, opacity: 0.85 }}>
                    📍 POSISI PAKET SEKARANG
                  </p>
                  <p style={{ margin: "0 0 2px", fontWeight: 700, fontSize: 16 }}>
                    {currentStep.currentLocation ?? currentStep.lokasi}
                  </p>
                  <p style={{ margin: 0, fontSize: 12, opacity: 0.85 }}>
                    {currentStep.title}
                  </p>
                </div>
              )}

              {/* Info singkat */}
              <div style={{ display: "flex", gap: 10, flexWrap: "wrap", marginBottom: 18 }}>
                {result.warehouse && <Chip label="No Resi"  val={result.warehouse.tracking_number} />}
                {result.warehouse && <Chip label="Barang"   val={result.warehouse.item_name || "-"} />}
                {result.shipment  && <Chip label="Tujuan"   val={result.shipment.destination_city} />}
                {result.delivery  && <Chip label="Kurir"    val={result.delivery.courier_name || "Auto-assigned"} />}
              </div>

              {/* Timeline gudang */}
              <p style={{ fontWeight: 700, fontSize: 13, color: "#374151", marginBottom: 12 }}>
                Riwayat Perjalanan Paket:
              </p>

              {timeline.map((step, i) => (
                <div key={i} style={{ display: "flex", gap: 14 }}>
                  {/* Ikon + garis */}
                  <div style={{ display: "flex", flexDirection: "column", alignItems: "center" }}>
                    <div style={{
                      width: 38, height: 38, borderRadius: "50%", flexShrink: 0,
                      background: step.current ? "#2563eb" : step.done ? "#eff6ff" : "#f1f5f9",
                      border: `2px solid ${step.current ? "#2563eb" : step.done ? "#93c5fd" : "#e2e8f0"}`,
                      display: "flex", alignItems: "center", justifyContent: "center",
                      fontSize: 18,
                      boxShadow: step.current ? "0 0 0 4px #bfdbfe" : "none",
                    }}>
                      {step.icon}
                    </div>
                    {i < timeline.length - 1 && (
                      <div style={{
                        width: 2, flexGrow: 1, minHeight: 20,
                        background: step.done ? "#bfdbfe" : "#e2e8f0",
                        margin: "4px 0",
                      }} />
                    )}
                  </div>

                  {/* Konten */}
                  <div style={{ paddingBottom: i < timeline.length - 1 ? 18 : 0, paddingTop: 6 }}>
                    <div style={{ display: "flex", alignItems: "center", gap: 8, flexWrap: "wrap" }}>
                      <span style={{
                        fontWeight: 700, fontSize: 13,
                        color: step.done ? "#111827" : "#9ca3af",
                      }}>
                        {step.title}
                      </span>
                      <Badge status={step.status} />
                      {step.current && (
                        <span style={{
                          background: "#fef9c3", color: "#854d0e",
                          fontSize: 10, fontWeight: 700,
                          padding: "2px 8px", borderRadius: 20,
                          border: "1px solid #fde047",
                        }}>
                          ● SEKARANG
                        </span>
                      )}
                    </div>
                    <p style={{ margin: "3px 0 1px", fontSize: 12, color: "#2563eb", fontWeight: 600 }}>
                      📍 {step.lokasi}
                    </p>
                    <p style={{ margin: 0, fontSize: 11, color: "#6b7280" }}>
                      🕐 {step.waktu}
                    </p>
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>

        {/* ── Tabel Semua Shipment ── */}
        <h2 style={{ fontSize: 16, marginBottom: 12, color: "#1a1a2e" }}>
          Semua Shipment
        </h2>

        {shipments.length === 0 ? (
          <p style={{ color: "#888" }}>Belum ada shipment.</p>
        ) : (
          <div style={{ overflowX: "auto" }}>
            <table style={{
              width: "100%", borderCollapse: "collapse", fontSize: 13,
              border: "1px solid #e5e7eb",
            }}>
              <thead style={{ background: "#c0392b", color: "white" }}>
                <tr>
                  {["ID","Tracking ID","No Resi","Asal","Tujuan","Lokasi Saat Ini","Status","ETA"].map(h => (
                    <th key={h} style={{ padding: "10px 12px", textAlign: "left", fontWeight: 600 }}>{h}</th>
                  ))}
                </tr>
              </thead>
              <tbody>
                {shipments.map((s) => (
                  <tr key={s.shipment_id} style={{
                    background: s.no_resi === searched ? "#eff6ff" : "white",
                    borderBottom: "1px solid #f1f5f9",
                  }}>
                    <td style={{ padding: "10px 12px" }}>{s.shipment_id}</td>
                    <td style={{ padding: "10px 12px" }}>{s.tracking_id}</td>
                    <td style={{ padding: "10px 12px", fontWeight: 600 }}>{s.no_resi}</td>
                    <td style={{ padding: "10px 12px" }}>{s.origin_city}</td>
                    <td style={{ padding: "10px 12px" }}>{s.destination_city}</td>
                    <td style={{ padding: "10px 12px" }}>{s.current_location}</td>
                    <td style={{ padding: "10px 12px" }}>
                      <Badge status={s.status} />
                    </td>
                    <td style={{ padding: "10px 12px" }}>
                      {s.eta ? new Date(s.eta).toLocaleString("id-ID") : "-"}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </>
  );
}

function Chip({ label, val }) {
  return (
    <div style={{
      background: "#f0f9ff", border: "1px solid #bae6fd",
      borderRadius: 8, padding: "4px 12px", fontSize: 12,
    }}>
      <span style={{ color: "#64748b" }}>{label}: </span>
      <span style={{ fontWeight: 700, color: "#0369a1" }}>{val}</span>
    </div>
  );
}
