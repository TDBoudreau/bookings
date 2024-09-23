function Prompt() {
    let toast = function (c) {
        const {msg = "", icon = "success", position = "top-end"} = c;

        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            position: position,
            icon: icon,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener("mouseenter", Swal.stopTimer);
                toast.addEventListener("mouseleave", Swal.resumeTimer);
            },
        });

        Toast.fire({});
    };

    let success = function (c) {
        const {msg = "", title = "", footer = ""} = c;

        Swal.fire({
            icon: "success",
            title: title,
            text: msg,
            footer: footer,
        });
    };

    let error = function (c) {
        const {msg = "", title = "", footer = ""} = c;

        Swal.fire({
            icon: "error",
            title: title,
            text: msg,
            footer: footer,
        });
    };

    async function custom(c) {
        const {
            msg = "",
            title = "",
            icon = "",
            showConfirmButton = true,
            showCancelButton = true,
            customClass = {
                confirmButton: "btn btn-primary mr-2",
                cancelButton: "btn btn-secondary",
            },
            buttonsStyling = false,
        } = c;

        const {value: result} = await Swal.fire({
            title: title,
            html: msg,
            icon: icon,
            // backdrop: false,
            focusConfirm: false,
            showConfirmButton: showConfirmButton,
            showCancelButton: showCancelButton,
            customClass: customClass,
            buttonsStyling: buttonsStyling,
            willOpen: () => {
                if (c.willOpen !== undefined) {
                    c.willOpen();
                }
            },
            didOpen: () => {
                if (c.didOpen !== undefined) {
                    c.didOpen();
                }
            },
            preConfirm: async () => {
                if (c.preConfirm !== undefined) {
                    // Ensure preConfirm is handled as a promise
                    const res = await c.preConfirm();
                    return res; // Return the result to SweetAlert's handling logic
                }
            },
        });

        if (result) {
            if (result.dismiss !== Swal.DismissReason.cancel) {
                if (result.value !== "") {
                    if (c.callback !== undefined) {
                        c.callback(result);
                    } else {
                        c.callback(false);
                    }
                }
            } else {
                c.callback(false);
            }
        }
    }

    return {
        toast: toast,
        success: success,
        error: error,
        custom: custom,
    };
}
