function roomAvailability(roomId, CSRFToken) {
    document
        .getElementById("check-availability-button")
        .addEventListener("click", function () {
            let html = `
                  <form id="check-availability-form" action="" method="post" novalidate class="needs-validation">
                      <div class="form-row">
                          <div class="col">
                              <div class="form-row" id="reservation-dates-modal">
                                  <div class="col">
                                      <input disabled required class="form-control" type="text" name="start" id="start" placeholder="Arrival" autocomplete="off">
                                  </div>
                                  <div class="col">
                                      <input disabled required class="form-control" type="text" name="end" id="end" placeholder="Departure" autocomplete="off">
                                  </div>
                              </div>
                          </div>
                      </div>
                  </form>
                  `;
            attention.custom({
                title: "Choose your dates",
                msg: html,

                willOpen: () => {
                    const elem = document.getElementById("reservation-dates-modal");
                    const rp = new DateRangePicker(elem, {
                        format: "yyyy-mm-dd",
                        minDate: new Date(),
                        showOnFocus: true,
                    });
                },
                didOpen: () => {
                    document.getElementById("start").removeAttribute("disabled");
                    document.getElementById("end").removeAttribute("disabled");
                },
                preConfirm: () => {
                    // Get the form element
                    const form = document.getElementById("check-availability-form");
                    // Trigger Bootstrap validation
                    form.classList.add("was-validated");
                    // Check form validity
                    if (form.checkValidity()) {
                        // If the form is valid, return the dates
                        return [
                            document.getElementById("start").value,
                            document.getElementById("end").value,
                        ];
                    } else {
                        // If the form is invalid, return false to prevent modal from closing
                        return false;
                    }
                },
                callback: function (result) {
                    let form = document.getElementById("check-availability-form");
                    let formData = new FormData(form);
                    formData.append("csrf_token", `${CSRFToken}`);
                    formData.append("room_id", `${roomId}`);
                    fetch("/search-availability-json", {
                        method: "post",
                        body: formData,
                    })
                        .then((response) => response.json())
                        .then((data) => {
                            if (data.ok) {
                                attention.custom({
                                    icon: "success",
                                    showConfirmButton: false,
                                    customClass: {
                                        confirmButton: "btn btn-primary mr-2",
                                        cancelButton: "btn btn-secondary"
                                    },
                                    msg: `<p>Room is available<p>
                                            <p>
                                                <a href="/book-room?id=${data.room_id}&s=${data.start_date}&e=${data.end_date}"
                                                    class="btn btn-primary">
                                                Book now!
                                            </a></p>`
                                })
                            } else {
                                attention.error({
                                    title: "No availability",
                                })
                            }
                        });
                },
            });
        });
}