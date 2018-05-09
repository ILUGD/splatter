#include "newform.h"
#include "ui_newform.h"

NewForm::NewForm(QWidget *parent) :
    QWidget(parent),
    ui(new Ui::NewForm)
{
    ui->setupUi(this);
    ui->dateEdit->setDate(QDate::currentDate());
    ui->fromTime->setTime(QTime::currentTime());
    ui->toTime->setTime(QTime::currentTime());
}

NewForm::~NewForm()
{
    delete ui;
}
