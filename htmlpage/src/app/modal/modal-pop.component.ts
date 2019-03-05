import { Component, OnInit } from '@angular/core';
import { BsModalRef } from 'ngx-bootstrap/modal';

@Component({
  selector: 'app-modal',
  templateUrl: './modal-pop.component.html',
  styleUrls: ['./modal-pop.component.css']
})

export class ModalComponent implements OnInit {
  constructor(public bsModalRef: BsModalRef) {
    console.log(this.bsModalRef.content);
  }
  public value: string = "确认删除吗";
  ngOnInit() {
  }
  title: string;//调用者给title进行了赋值
  isCancel: boolean = true;
  btnCloseClick() {
    //this.bsModalRef.content = "===";
    this.bsModalRef.hide();
  }

  btnConfirmClick() {
    this.isCancel = false;
    this.bsModalRef.hide();
  }
}