package main

import (
	    "encoding/csv"
	        "fmt"
		    "log"
		        "os"
			    "time"

			        "github.com/go-resty/resty/v2"
			)

			type Transaction struct {
				    Date        time.Time `csv:"Date"`
				        SentAmount  float64   `csv:"Sent Amount"`
					    SentCurrency string   `csv:"Sent Currency"`
					        ReceivedAmount float64 `csv:"Received Amount"`
						    ReceivedCurrency string `csv:"Received Currency"`
						        FeeAmount    float64   `csv:"Fee Amount"`
							    FeeCurrency  string    `csv:"Fee Currency"`
							        Tag          string    `csv:"Tag"`
							}

							func main() {
								    client := resty.New()

								        // Replace with your Ronin wallet API endpoint
									    resp, err := client.R().Get("https://api.roninchain.com/wallet/transactions")
									        if err != nil {
											        log.Fatalf("Error fetching data: %v", err)
												    }

												        // Parse the response (this is a placeholder, adjust according to the actual API response)
													    var transactions []Transaction
													        // Assume the response is in JSON and unmarshal it into transactions slice
														    // json.Unmarshal(resp.Body(), &transactions)

														        // Create a CSV file
															    file, err := os.Create("koinly_transactions.csv")
															        if err != nil {
																	        log.Fatalf("Could not create file: %v", err)
																		    }
																		        defer file.Close()

																			    writer := csv.NewWriter(file)
																			        defer writer.Flush()

																				    // Write CSV header
																				        writer.Write([]string{"Date", "Sent Amount", "Sent Currency", "Received Amount", "Received Currency", "Fee Amount", "Fee Currency", "Tag"})

																					    // Write transactions to CSV
																					        for _, tx := range transactions {
																							        record := []string{
																									            tx.Date.Format("2006-01-02 15:04:05"),
																										                fmt.Sprintf("%f", tx.SentAmount),
																												            tx.SentCurrency,
																													                fmt.Sprintf("%f", tx.ReceivedAmount),
																															            tx.ReceivedCurrency,
																																                fmt.Sprintf("%f", tx.FeeAmount),
																																		            tx.FeeCurrency,
																																			                tx.Tag,
																																					        }
																																						        writer.Write(record)
																																							    }

																																							        fmt.Println("CSV file created successfully")
																																							}

