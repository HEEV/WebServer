//run when the page loads
$(function() {
/*  $('#GetExcel').on('click', function(e) {
    e.preventDefault();
    $('#GetExcel').fadeOut(300);

    var params = {};

    $.ajax({
      url: 'API/ajaxGetRace.php',
      type: 'post',
      data: params,
      success: function(data, status) {
        if (status == 200) {
          //make the file downloadable here
        }
      },
      error: function(xhr, desc, err) {
        console.log(xhr);
        console.log('Details: ' + desc + '\nError: ' + err);
      }
    }); //end ajax call
  });*/

  $.ajax({
    url: 'API/ajaxGetRaceIds.php',
    type: 'get',
    success: function(data, status) {
      if(status == 200) {
        var dropdown = $('#RaceId');
        dropdown.empty();
        dropdown.append($('<option></option>').attr('value', '').text('Please Select'));
        $.each(data, function(value, key) {
          dropdown.append($('<option></option>').attr('value', value).text(key));
        });
      }
    },
    error: function(xhr, desc, err) {
      console.log(xhr);
      console.log('Details: ' + desc + '\nError: ' + err);
    }
  }); //end ajax call

  $('#GetExcel').disable();

  $('#RaceId').on('change', function(e) {
    if (e.text == 'Please Select') {
      $('#GetExcel').disable();
    }
    else {
      $('#GetExcel').enable();
    }
  });
});
