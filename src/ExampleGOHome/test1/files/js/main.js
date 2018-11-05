$('#myButton').on('click', function(e){
    var name = $('.name').val();
    var lastName = $('.lastName').val();
    
    $.ajax({
        url: '/ajax/users/add',
        data: {
            name: name,
            lastName: lastName
        },
        type: 'POST',
        timeout: 15000,
        success: function (result) {
            var user = result.data.user;
            $('#maintable').children('tbody').append(`
            <tr data-id="`+user.id+`">
            <td>`+user.name+`</td>
            <td>`+user.lastName+`</td>
            <td><button class="remove">X</button></td>
            </tr>`);
        },
        error: function (result) {
            console.log(result);
        }
    });
})

$('body').on('click', '.remove', function(e){
    console.log("target:", $(e.target));
    var id = $(e.target).parent().parent().data("id");
    var name = $('.name').val();
    var lastName = $('.lastName').val();
    
    console.log("ID:", id);
    $.ajax({
        url: '/ajax/users/remove',
        data: {
            id: id
        },
        type: 'POST',
        timeout: 15000,
        success: function (result) {
            var id = result.data.id;
            $('[data-id="'+id+'"]').remove();
        },
        error: function (result) {
            console.log(result);
        }
    });
})