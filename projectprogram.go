package main
import "fmt"
const NMAX = 30
type Latihan struct {
	NamaGerakan string
	Otot        string
	Set         int
	Reps        int
}
var NewbieUpper [NMAX]Latihan
var nUpper int
var NewbieMid [NMAX]Latihan
var nMid int
var NewbieLower [NMAX]Latihan
var nLower int

func isiDataNewbieFix() {
	NewbieUpper[0] = Latihan{"Bench_Press", "Dada", 4, 12}
	NewbieUpper[1] = Latihan{"Pull_Up", "Punggung", 4, 8}
	NewbieUpper[2] = Latihan{"Bicep_Curl", "Lengan", 3, 15}
	NewbieUpper[3] = Latihan{"Shoulder_Press", "Bahu", 3, 10}
	nUpper = 4

	NewbieMid[0] = Latihan{"Crunch", "Perut", 3, 20}
	NewbieMid[1] = Latihan{"Plank", "Core", 3, 60}
	NewbieMid[2] = Latihan{"Russian_Twist", "Samping", 3, 15}
	nMid = 3

	NewbieLower[0] = Latihan{"Squat", "Paha", 4, 12}
	NewbieLower[1] = Latihan{"Deadlift", "Hamstring", 3, 8}
	NewbieLower[2] = Latihan{"Calf_Raise", "Betis", 4, 20}
	NewbieLower[3] = Latihan{"Leg_Press", "Paha", 3, 10}
	NewbieLower[4] = Latihan{"Lunges", "Paha", 3, 12}
	nLower = 5
}

func cetakLatihan(A [NMAX]Latihan, n int, judul string) {
	fmt.Printf("\n--- %s ---\n", judul)
	if n == 0 {
		fmt.Println("Belum ada data latihan.")
	} else {
		for i := 0; i < n; i++ {
			fmt.Printf("%d. %s (%s) - %d Set x %d Reps\n", i+1, A[i].NamaGerakan, A[i].Otot, A[i].Set, A[i].Reps)
		}
	}
}
func selectionSortLatihan(A [NMAX]Latihan, n int, isAscending bool) [NMAX]Latihan {
	for i := 0; i < n-1; i++ {
		idxEkstrim := i
		for j := i + 1; j < n; j++ {
			if isAscending {
				if A[j].Reps < A[idxEkstrim].Reps {
					idxEkstrim = j
				}
			} else {
				if A[j].Reps > A[idxEkstrim].Reps {
					idxEkstrim = j
				}
			}
		}
		temp := A[idxEkstrim]
		A[idxEkstrim] = A[i]
		A[i] = temp
	}
	return A
}
func insertionSortLatihan(A [NMAX]Latihan, n int, isAscending bool) [NMAX]Latihan {
	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1
		if isAscending {
			for j >= 0 && A[j].NamaGerakan > key.NamaGerakan {
				A[j+1] = A[j]
				j--
			}
		} else {
			for j >= 0 && A[j].NamaGerakan < key.NamaGerakan {
				A[j+1] = A[j]
				j--
			}
		}
		A[j+1] = key
	}
	return A
}
func sequentialSearch(A [NMAX]Latihan, n int, x string) int {
	ketemu := false
	k := 0
	for !ketemu && k < n {
		ketemu = A[k].NamaGerakan == x
		if !ketemu {
			k = k + 1
		}
	}
	if ketemu {
		return k
	}
	return -1
}
func binarySearch(A [NMAX]Latihan, n int, x string) int {
	kiri := 0
	kanan := n - 1
	ketemu := false
	mid := 0

	for kiri <= kanan && !ketemu {
		mid = (kiri + kanan) / 2
		if A[mid].NamaGerakan == x {
			ketemu = true
		} else if A[mid].NamaGerakan < x {
			kiri = mid + 1
		} else {
			kanan = mid - 1
		}
	}

	if ketemu {
		return mid
	}
	return -1
}

func tambahLatihanPro(A [NMAX]Latihan, n int) ([NMAX]Latihan, int) {
	if n < NMAX {
		fmt.Println("Masukkan (NamaGerakan Otot Set Reps):")
		fmt.Scan(&A[n].NamaGerakan)
		fmt.Scan(&A[n].Otot)
		fmt.Scan(&A[n].Set)
		fmt.Scan(&A[n].Reps)
		n = n + 1
		fmt.Println("Data berhasil ditambahkan!")
	} else {
		fmt.Println("Kapasitas latihan penuh.")
	}
	return A, n
}

func hapusLatihanPro(A [NMAX]Latihan, n int, namaHapus string) ([NMAX]Latihan, int) {
	idx := sequentialSearch(A, n, namaHapus) // Manfaatkan sequential search
	if idx != -1 {
		for i := idx; i < n-1; i++ {
			A[i] = A[i+1]
		}
		n = n - 1
		fmt.Println(">> Data berhasil dihapus.")
	} else {
		fmt.Println(">> Data tidak ditemukan.")
	}
	return A, n
}
func editLatihanPro(A [NMAX]Latihan, n int) [NMAX]Latihan {
	if n == 0 {
		fmt.Println(">> Data latihan masih kosong!")
		return A
	}

	var namaEdit string
	fmt.Print("Masukkan Nama Gerakan yang ingin diedit: ")
	fmt.Scan(&namaEdit)

	idx := sequentialSearch(A, n, namaEdit)

	if idx != -1 {
		fmt.Printf("Data saat ini: %s | %d Set | %d Reps\n", A[idx].NamaGerakan, A[idx].Set, A[idx].Reps)
		fmt.Print("Masukkan Set baru : ")
		fmt.Scan(&A[idx].Set)
		fmt.Print("Masukkan Reps baru: ")
		fmt.Scan(&A[idx].Reps)
		fmt.Println(">> Data berhasil diupdate!")
	} else {
		fmt.Println(">> Gerakan tidak ditemukan!")
	}
	return A
}

func cetakReferensiExercise() {
	fmt.Println("\n========================================")
	fmt.Println("REFERENSI EXERCISE")
	fmt.Println("========================================")
	fmt.Println("[UPPER BODY]")
	for i := 0; i < nUpper; i++ {
		fmt.Printf("  - %s (%s)\n", NewbieUpper[i].NamaGerakan, NewbieUpper[i].Otot)
	}
	fmt.Println("[MID BODY]")
	for i := 0; i < nMid; i++ {
		fmt.Printf("  - %s (%s)\n", NewbieMid[i].NamaGerakan, NewbieMid[i].Otot)
	}
	fmt.Println("[LOWER BODY]")
	for i := 0; i < nLower; i++ {
		fmt.Printf("  - %s (%s)\n", NewbieLower[i].NamaGerakan, NewbieLower[i].Otot)
	}
	fmt.Println("========================================")
}

func main() {
	isiDataNewbieFix()

	var ProSenin, ProSelasa, ProRabu [NMAX]Latihan
	var nSenin, nSelasa, nRabu int

	menuUtamaJalan := true
	var pilihanUtama int

	for menuUtamaJalan {
		fmt.Println("\n=== SELAMAT DATANG DI GYM MANJAPAHIT PARIZ ===")
		fmt.Println("1. Mode NOOB (antum belum pro jadinya jangan gegabah)")
		fmt.Println("2. Mode PRO GYM (njier sepuh coy)")
		fmt.Println("3. Keluar Aplikasi")
		fmt.Print("Pilih opsi (1/2/3): ")
		fmt.Scan(&pilihanUtama)

		if pilihanUtama == 1 {
			menuNewbieJalan := true
			var pilihanNewbie int

			for menuNewbieJalan {
				fmt.Println("\n- PILIHAN TARGET MUSCLE (NEWBIE) -")
				fmt.Println("1. Paket Upper Body (Push/Pull)")
				fmt.Println("2. Paket Mid Body (Core/Abs)")
				fmt.Println("3. Paket Lower Body (Legs)")
				fmt.Println("4. Kembali ke Menu Utama")
				fmt.Print("Pilih Target Latihan (1/2/3/4): ")
				fmt.Scan(&pilihanNewbie)

				if pilihanNewbie == 1 || pilihanNewbie == 2 || pilihanNewbie == 3 {
					var DataPilih [NMAX]Latihan
					var nPilih int
					var judul string

					if pilihanNewbie == 1 {
						DataPilih = NewbieUpper
						nPilih = nUpper
						judul = "PAKET UPPER BODY"
					} else if pilihanNewbie == 2 {
						DataPilih = NewbieMid
						nPilih = nMid
						judul = "PAKET MID BODY"
					} else if pilihanNewbie == 3 {
						DataPilih = NewbieLower
						nPilih = nLower
						judul = "PAKET LOWER BODY"
					}

					cetakLatihan(DataPilih, nPilih, judul)
					aksiNewbieJalan := true
					var aksi int
					for aksiNewbieJalan {
						fmt.Println("\nAksi pada paket ini:")
						fmt.Println("1. Sort by Reps - Selection ASC")
						fmt.Println("2. Sort by Reps - Selection DESC")
						fmt.Println("3. Cari Gerakan (Sequential)")
						fmt.Println("4. Selesai (Ganti Paket)")
						fmt.Print("Pilih Aksi: ")
						fmt.Scan(&aksi)

						if aksi == 1 {
							DataPilih = selectionSortLatihan(DataPilih, nPilih, true)
							cetakLatihan(DataPilih, nPilih, judul+" (Sorted Reps ASC)")
						} else if aksi == 2 {
							DataPilih = selectionSortLatihan(DataPilih, nPilih, false)
							cetakLatihan(DataPilih, nPilih, judul+" (Sorted Reps DESC)")
						} else if aksi == 3 {
							var cari string
							fmt.Print("Masukkan Nama Gerakan yang dicari: ")
							fmt.Scan(&cari)
							idx := sequentialSearch(DataPilih, nPilih, cari)
							if idx != -1 {
								fmt.Printf(">> KETEMU: %s ada di urutan ke-%d\n", DataPilih[idx].NamaGerakan, idx+1)
							} else {
								fmt.Println(">> Gerakan tidak ada di paket ini.")
							}
						} else if aksi == 4 {
							aksiNewbieJalan = false
						} else {
							fmt.Println("Aksi tidak valid.")
						}
					}
				} else if pilihanNewbie == 4 {
					menuNewbieJalan = false
				} else {
					fmt.Println("Pilihan salah.")
				}
			}

		} else if pilihanUtama == 2 {
			menuProJalan := true
			var pilihanPro int

			for menuProJalan {
				fmt.Println("\n--- MENU PRO GYM (CUSTOM HARIAN) ---")
				fmt.Println("1. Kelola Latihan (Senin/Selasa/Rabu)")
				fmt.Println("2. Kembali ke Menu Utama")
				fmt.Print("Pilih: ")
				fmt.Scan(&pilihanPro)

				if pilihanPro == 1 {
					var hari string
					fmt.Print("Input Hari (Senin/Selasa/Rabu): ")
					fmt.Scan(&hari)

					if hari == "Senin" || hari == "Selasa" || hari == "Rabu" {
						cetakReferensiExercise()
						subProJalan := true
						for subProJalan {
							fmt.Printf("\n--- OPSI CUSTOM: %s ---\n", hari)
							fmt.Println("1. Tambah Exercise")
							fmt.Println("2. Hapus Exercise")
							fmt.Println("3. Edit Exercise (Set & Reps)")
							fmt.Println("4. Cari Exercise (Sequential/Binary)")
							fmt.Println("5. Lihat & Urutkan Exercise")
							fmt.Println("6. Selesai (Kembali)")
							fmt.Print("Pilih Opsi: ")
							fmt.Scan(&pilihanPro)
							var dataTemp [NMAX]Latihan
							var nTemp int
							if hari == "Senin" {
								dataTemp = ProSenin
								nTemp = nSenin
							} else if hari == "Selasa" {
								dataTemp = ProSelasa
								nTemp = nSelasa
							} else {
								dataTemp = ProRabu
								nTemp = nRabu
							}

							if pilihanPro == 1 {
								dataTemp, nTemp = tambahLatihanPro(dataTemp, nTemp)
							} else if pilihanPro == 2 {
								var namaHapus string
								fmt.Print("Nama Exercise yang dihapus: ")
								fmt.Scan(&namaHapus)
								dataTemp, nTemp = hapusLatihanPro(dataTemp, nTemp, namaHapus)
							} else if pilihanPro == 3 {
								dataTemp = editLatihanPro(dataTemp, nTemp)
							} else if pilihanPro == 4 {
								var cari string
								var metodeCari int
								fmt.Print("Masukkan Nama Gerakan yang dicari: ")
								fmt.Scan(&cari)
								fmt.Println("Metode Pencarian:")
								fmt.Println("1. Sequential Search")
								fmt.Println("2. Binary Search")
								fmt.Print("Pilih: ")
								fmt.Scan(&metodeCari)

								if metodeCari == 1 {
									idx := sequentialSearch(dataTemp, nTemp, cari)
									if idx != -1 {
										fmt.Printf(">> (Sequential) KETEMU: %s dengan %d Set, %d Reps\n", dataTemp[idx].NamaGerakan, dataTemp[idx].Set, dataTemp[idx].Reps)
									} else {
										fmt.Println(">> Data tidak ditemukan!")
									}
								} else if metodeCari == 2 {
									dataTemp = insertionSortLatihan(dataTemp, nTemp, true)
									fmt.Println(">> Info: Data otomatis diurutkan (A-Z) untuk Binary Search.")
									idx := binarySearch(dataTemp, nTemp, cari)
									if idx != -1 {
										fmt.Printf(">> (Binary) KETEMU: %s dengan %d Set, %d Reps\n", dataTemp[idx].NamaGerakan, dataTemp[idx].Set, dataTemp[idx].Reps)
									} else {
										fmt.Println(">> Data tidak ditemukan!")
									}
								}
							} else if pilihanPro == 5 {
								var metodeSort int
								fmt.Println("1. Tampilkan Biasa")
								fmt.Println("2. Selection Sort (Reps ASC)")
								fmt.Println("3. Insertion Sort (Nama A-Z)")
								fmt.Print("Pilih: ")
								fmt.Scan(&metodeSort)

								if metodeSort == 2 {
									dataTemp = selectionSortLatihan(dataTemp, nTemp, true)
								} else if metodeSort == 3 {
									dataTemp = insertionSortLatihan(dataTemp, nTemp, true)
								}
								cetakLatihan(dataTemp, nTemp, "DAFTAR LATIHAN "+hari)
							} else if pilihanPro == 6 {
								subProJalan = false
							} else {
								fmt.Println("Pilihan salah.")
							}
							if hari == "Senin" {
								ProSenin = dataTemp
								nSenin = nTemp
							} else if hari == "Selasa" {
								ProSelasa = dataTemp
								nSelasa = nTemp
							} else {
								ProRabu = dataTemp
								nRabu = nTemp
							}
						}
					} else {
						fmt.Println("Hari tidak valid.")
					}

				} else if pilihanPro == 2 {
					menuProJalan = false
				} else {
					fmt.Println("Pilihan salah.")
				}
			}

		} else if pilihanUtama == 3 {
			menuUtamaJalan = false
			fmt.Println("Program selesai. Keep grinding!")
		} else {
			fmt.Println("Pilihan tidak ada.")
		}
	}
}