using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace InputKue
{
    public class Kue
    {
        private String kodeKue;
        private String namaKue;
        private int jumlah;
        private int harga;

        public Kue(string kode, string nama, int jmlh, int harga)
        {
            this.kodeKue = kode;
            this.namaKue = nama;
            this.jumlah = jmlh;
            this.harga = harga;
        }

        public void setKodeKue(string kode)
        {
            this.kodeKue = kode;
        }

        public String getKodeKue()
        {
            return kodeKue;
        }
        public void setNamaKue(string nama)
        {
            this.namaKue = nama;
        }

        public String getNamaKue()
        {
            return namaKue;
        }
        public void setJumlah(int jmlh)
        {
            this.jumlah = jmlh;
        }

        public int getJumlah()
        {
            return jumlah;
        }
        public void setHarga(int harga)
        {
            this.harga = harga;
        }

        public int getHarga()
        {
            return harga;
        }
    }
}
