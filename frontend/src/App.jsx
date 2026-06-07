import { useState } from "react";
import "./App.css";

function App() {
  const [label, setLabel] = useState("");
  const [nilai, setNilai] = useState("");
  const [hasil, setHasil] = useState(null);

  function formatRupiah(angka) {
    if (!angka) return "";
    return new Intl.NumberFormat("id-ID").format(angka);
  }

  const handleNilaiChange = (e) => {
    const raw = e.target.value.replace(/\D/g, "");
    setNilai(raw);
  };

  const hitung = async () => {
    const response = await fetch("http://localhost:8080/konversi", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        label,
        nilai: Number(nilai),
      }),
    });

    const data = await response.json();
    setHasil(data);
  };

  const reset = () => {
    setLabel("");
    setNilai("");
    setHasil(null);
  };

  return (
    <div className="app-container">
      <div className="card">
        <h1 className="title">MBG Converter</h1>

        <div className="input-group">
          <label className="label">Apa yang ingin di konversi?</label>
          <input
            type="text"
            className="input-field"
            value={label}
            onChange={(e) => setLabel(e.target.value)}
            placeholder="Maukan label"
          />
        </div>

        <div className="input-group">
          <label className="label">Nilai</label>
          <input
            type="text"
            className="input-field"
            value={formatRupiah(nilai)}
            onChange={handleNilaiChange}
            placeholder="Masukan angka"
          />
        </div>

        <div className="button-container">
          <button className="action-button" onClick={hitung}>
            Hitung
          </button>
          <button className="action-button" onClick={reset}>
            Reset
          </button>
        </div>

        {hasil && (
          <div className="result-container">
            <div className="result-label">Result</div>
            <div className="result-details">
              <div className="result-item">
                <span className="result-label-text">{hasil.label}</span>
                <span className="result-value">Rp{formatRupiah(nilai)}</span>
              </div>
              <div className="result-divider"></div>
              <div className="result-item">
                <span className="result-label-text">Equivalent</span>
                <span className="result-value-large">{hasil.hasil.toFixed(3)}</span>
                <span className="result-unit">{hasil.satuan} MBG</span>
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

export default App;
