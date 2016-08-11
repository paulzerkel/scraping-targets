'use strict';

function formatCurrency(value) {
  let tempValue = parseInt(value) / 100;
  return "$" + tempValue.toFixed(2);
}

function formatWeight(value) {
  return value + " lbs";
}

function Product(data) {
  this.id = ko.observable(data.id);
  this.name = ko.observable(data.name);
  this.description = ko.observable(data.description);
  this.price = ko.observable(data.price);
  this.weight = ko.observable(data.weight);
}

function ProductViewModel() {
  var self = this;
  self.products = ko.observableArray([])

  $.getJSON("/product-data", function(rawProducts) {
    let productObj = $.map(rawProducts, function(item) { return new Product(item) });
    self.products(productObj);
  });
}

$(function() {
  ko.applyBindings(new ProductViewModel());
});
